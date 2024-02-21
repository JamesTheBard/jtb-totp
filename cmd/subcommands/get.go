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

	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/pquerna/otp/totp"
)

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Generate TOTP code from key",
	Args:  cobra.ExactArgs(1),
	Run:   getCommand,
}

func getCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))

	name, value, _ := getTotpKey(args[0], &data)
	t := time.Now()
	code, err := totp.GenerateCode(value, t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s -> %s\n", name, code)
}

func getTotpKey(search string, dataMap *map[string]string) (string, string, int) {
	score, name := -1, ""
	for n := range *dataMap {
		if v := fuzzy.RankMatchFold(search, n); v > score {
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
