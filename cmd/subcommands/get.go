package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"

	fuzzy "github.com/paul-mannino/go-fuzzywuzzy"
	"github.com/pquerna/otp/totp"
)

func init() {
	var exact bool
	getCmd.Flags().BoolVarP(&exact, "exact", "e", false, "match the key name exactly")

	cmd.RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get [key name]",
	Short: "Generate TOTP code from key",
	Args:  cobra.ExactArgs(1),
	Run:   getCommand,
}

func getCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	err := keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))
	if err != nil {
		fmt.Printf("Could not open/process the keystore: %s\n", err)
		os.Exit(1)
	}

	exact, _ := cmd.Flags().GetBool("exact")
	name, value, score := getTotpKey(args[0], &data, exact)
	t := time.Now()
	code, err := totp.GenerateCode(value, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("[%d%%] %s: %s\n", score, name, code)
}

func getTotpKey(search string, dataMap *map[string]string, exact bool) (string, string, int) {
	if exact {
		for n := range *dataMap {
			if search == n {
				return n, (*dataMap)[n], 100
			}
		}
		fmt.Println("Could not find a key!")
		os.Exit(1)
	}

	score, name := -1, ""
	for n := range *dataMap {
		if v := fuzzy.PartialRatio(search, n); v > score {
			score = v
			name = n
		}
	}

	if score < 0 {
		fmt.Println("Could not find a key!")
		os.Exit(1)
	}

	return name, (*dataMap)[name], score
}
