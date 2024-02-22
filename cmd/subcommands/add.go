package cmd

import (
	"fmt"
	"os"
	"strings"

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
	err := keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))
	if err != nil {
		fmt.Printf("Could not open/process the keystore: %s\n", err)
		os.Exit(1)
	}

	key, secret := strings.TrimSpace(args[0]), strings.TrimSpace(args[1])
	data[key] = secret
	yamlData := keystore.DumpYaml(&data)

	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Printf("Updated keystore with new/changed data.\n")
	fmt.Printf("Added key '%s' to keystore successfully!\n", args[0])
}
