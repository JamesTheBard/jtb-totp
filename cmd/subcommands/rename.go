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
	cmd.RootCmd.AddCommand(renameCmd)
}

var renameCmd = &cobra.Command{
	Use:   "rename [key name] [new name]",
	Short: "Rename key in keystore",
	Args:  cobra.ExactArgs(2),
	Run:   renameCommand,
}

func renameCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	err := keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))
	if err != nil {
		fmt.Printf("Could not open/process the keystore: %s\n", err)
		os.Exit(1)
	}

	old_key := strings.TrimSpace(args[0])
	new_key := strings.TrimSpace(args[1])

	if _, exists := data[new_key]; exists {
		fmt.Printf("The new key name '%s' already exists!  The new key name must not exist in the keystore.\n", new_key)
		os.Exit(1)
	}

	if secret, exists := data[old_key]; exists {
		delete(data, old_key)
		data[new_key] = secret
		fmt.Printf("Renamed '%s' to '%s'.\n", old_key, new_key)
	} else {
		fmt.Printf("Key '%s' does not exist!\n", old_key)
		os.Exit(1)
	}

	yamlData := keystore.DumpYaml(&data)

	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Printf("Updated keystore with new/changed data.\n")
}
