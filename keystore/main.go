package keystore

import (
	"fmt"
	"log"
	"os"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"gopkg.in/yaml.v2"
)

type TotpKey struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

func LoadYaml(filename string, dataMap *map[string]string) {
	f, err := os.ReadFile("keys.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var data []TotpKey
	if err := yaml.Unmarshal(f, &data); err != nil {
		log.Fatal(err)
	}

	for _, k := range data {
		(*dataMap)[k.Name] = k.Key
	}
}

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
	fmt.Println("Written!")
}
