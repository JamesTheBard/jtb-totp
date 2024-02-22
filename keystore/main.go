package keystore

import (
	"encoding/json"
	"fmt"
	"jamesthebard/jtb-totp/config"
	"log"
	"os"
	"strings"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"gopkg.in/yaml.v2"
)

type TotpKey struct {
	Name string `yaml:"name" json:"name"`
	Key  string `yaml:"key" json:"key"`
}

func LoadEncryptedYaml(filename string, dataMap *map[string]string, password []byte) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	encrypted := crypto.NewPGPMessage(f)
	decrypted, err := crypto.DecryptMessageWithPassword(encrypted, password)
	if err != nil {
		return err
	}

	var data []TotpKey
	if err := yaml.Unmarshal(decrypted.Data, &data); err != nil {
		return err
	}

	for _, k := range data {
		(*dataMap)[k.Name] = k.Key
	}

	return nil
}

func DumpYaml(data *map[string]string) []byte {
	yamlData := []TotpKey{}
	for k, v := range *data {
		yamlData = append(yamlData, TotpKey{
			Name: k,
			Key:  v,
		})
	}
	out, _ := yaml.Marshal(yamlData)
	return out
}

func EncryptKeystore(filename string, data []byte, password []byte) {
	message := crypto.NewPlainMessage(data)
	encrypted, err := crypto.EncryptMessageWithPassword(message, password)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(filename, encrypted.Data, 0600)
}

func ExportKeystore(filename string, stdout bool) {
	data := map[string]string{}
	var exists bool
	config.Password, exists = os.LookupEnv(config.PasswdEnvVar)
	if !exists {
		fmt.Println("In order to export the keystore, you must explicitly set the password via environment variable!")
		os.Exit(1)
	}

	LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))

	yamlData := DumpYaml(&data)
	if stdout {
		fmt.Println(string(yamlData))
	} else {
		err := os.WriteFile(filename, yamlData, 0400)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ImportKeystore(filename string, overwrite bool) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var data []TotpKey

	if err := yaml.Unmarshal(f, &data); err == nil && len(data) > 0 {
		ProcessImportData(&f, &data, overwrite)
		fmt.Printf("Loaded %d keys from YAML file '%s'.\n", len(data), filename)
		return
	}

	if err := json.Unmarshal(f, &data); err == nil && len(data) > 0 {
		ProcessImportData(&f, &data, overwrite)
		fmt.Printf("Loaded %d keys from JSON file '%s'.\n", len(data), filename)
		return
	}

	fmt.Printf("Cannot import '%s', either the file is corrupt, ill-formatted, or there are no keys in the file!\n", filename)
}

func ProcessImportData(fileData *[]byte, data *[]TotpKey, overwrite bool) {
	currentData := make(map[string]string)
	err := LoadEncryptedYaml(config.KeystoreFile, &currentData, []byte(config.Password))
	if err != nil {
		fmt.Printf("Could not open/process the keystore: %s\n", err)
		os.Exit(1)
	}

	for _, k := range *data {
		name := strings.TrimSpace(k.Name)
		if _, ok := currentData[name]; ok && !overwrite {
			continue
		}
		currentData[name] = k.Key
		fmt.Printf("Imported key '%s'...\n", name)
	}

	yamlData := DumpYaml(&currentData)
	EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
}
