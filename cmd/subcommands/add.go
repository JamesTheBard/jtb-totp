package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var KeyName string
	var KeyValue string

	// Add Command
	addCmd.Flags().StringVarP(&KeyName, "key-name", "n", "", "name of the key (required)")
	addCmd.Flags().StringVarP(&KeyValue, "key-value", "v", "", "value of the key (required)")
	addCmd.MarkFlagRequired("key-name")
	addCmd.MarkFlagRequired("key-value")

	cmd.RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add key to keystore",
	Run:   addCommand,
}

func addCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))

	key, _ := cmd.Flags().GetString("key-name")
	value, _ := cmd.Flags().GetString("key-value")

	data[key] = value
	yamlData := keystore.DumpYaml(&data)

	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Printf("Updated keystore with new/changed data.\n")
	fmt.Printf("Added key '%s' to keystore successfully!\n", key)
}
