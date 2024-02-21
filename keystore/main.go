package keystore

import (
	"encoding/json"
	"fmt"
	"jamesthebard/jtb-totp/config"
	"log"
	"os"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"gopkg.in/yaml.v2"
)

type TotpKey struct {
	Name string `yaml:"name" json:"name"`
	Key  string `yaml:"key" json:"key"`
}

// func LoadYaml(filename string, dataMap *map[string]string) {
// 	f, err := os.ReadFile("keys.yaml")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var data []TotpKey
// 	if err := yaml.Unmarshal(f, &data); err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, k := range data {
// 		(*dataMap)[k.Name] = k.Key
// 	}
// }

func LoadEncryptedYaml(filename string, dataMap *map[string]string, password []byte) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	encrypted := crypto.NewPGPMessage(f)
	decrypted, err := crypto.DecryptMessageWithPassword(encrypted, password)
	if err != nil {
		log.Fatal(err)
	}

	var data []TotpKey
	if err := yaml.Unmarshal(decrypted.Data, &data); err != nil {
		log.Fatal(err)
	}

	for _, k := range data {
		(*dataMap)[k.Name] = k.Key
	}
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

func ImportKeystore(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var data []TotpKey

	if err := yaml.Unmarshal(f, &data); err == nil && len(data) > 0 {
		ProcessImportData(&f, &data)
		fmt.Printf("Loaded %d keys from YAML file '%s'.\n", len(data), filename)
		return
	}

	if err := json.Unmarshal(f, &data); err == nil && len(data) > 0 {
		ProcessImportData(&f, &data)
		fmt.Printf("Loaded %d keys from JSON file '%s'.\n", len(data), filename)
		return
	}

	fmt.Printf("Cannot import '%s', either the file is corrupt, ill-formatted, or there are no keys in the file!\n", filename)
}

func ProcessImportData(fileData *[]byte, data *[]TotpKey) {
	keyMap := map[string]string{}
	for _, k := range *data {
		keyMap[k.Name] = k.Key
		fmt.Printf("Imported key '%s'...\n", k.Name)
	}
	yamlData := DumpYaml(&keyMap)
	EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
}
