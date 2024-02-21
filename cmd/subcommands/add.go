package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	cmd.RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add [key name] [key secret]",
	Short: "Add key to keystore",
	Args:  cobra.ExactArgs(2),
	Run:   addCommand,
}

func addCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))

	data[args[0]] = args[1]
	yamlData := keystore.DumpYaml(&data)

	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Printf("Updated keystore with new/changed data.\n")
	fmt.Printf("Added key '%s' to keystore successfully!\n", args[0])
}
