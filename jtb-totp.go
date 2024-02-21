package main

import (
	"jamesthebard/jtb-totp/cmd"
	_ "jamesthebard/jtb-totp/cmd/subcommands"
)

type TotpKey struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

type TotpKeyStore struct {
	KeyStore []TotpKey
}

func main() {
	cmd.Execute()

	// keysearch := os.Args[1]
	// var data map[string]string
	// encryptFile("keys.yaml", []byte("1234"))
	// parseYaml("keys.yaml.enc", &data)

	// name, key, _ := getTotpKey(keysearch, &data)

	// s, _ := totp.GenerateCode(key, time.Now())
	// fmt.Printf("%s: %s\n", name, s)
}

// func encryptFile(filename string, password []byte) {
// 	f, err := os.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	message := crypto.NewPlainMessage(f)
// 	encrypted, err := crypto.EncryptMessageWithPassword(message, []byte("1234"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	save := "keys.yaml.enc"
// 	os.WriteFile(save, encrypted.Data, 0400)
// }

// func getTotpKey(search string, dataMap *map[string]string) (string, string, int) {
// 	score, name := -1, ""
// 	for n := range *dataMap {
// 		if v := fuzzy.RankMatchFold(search, n); v > score {
// 			score = v
// 			name = n
// 		}
// 	}

// 	if score < 0 {
// 		fmt.Println("Could not find a key!")
// 		os.Exit(1)
// 	}

// 	return name, (*dataMap)[name], score
// }
