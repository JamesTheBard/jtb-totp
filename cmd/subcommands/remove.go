package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	cmd.RootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove [key name]",
	Short: "Remove key from keystore",
	Args:  cobra.ExactArgs(1),
	Run:   removeCommand,
}

func removeCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))
	key := strings.TrimSpace(args[0])

	_, ok := data[key]
	delete(data, key)
	yamlData := keystore.DumpYaml(&data)
	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	if ok {
		fmt.Printf("Deleted key '%s' from the keystore.\n", args[0])
	} else {
		fmt.Printf("Could not find key '%s' in the keystore.\n", args[0])
	}
}
