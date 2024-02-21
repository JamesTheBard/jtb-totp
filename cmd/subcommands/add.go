package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var KeyName string
	var KeyValue string

	// Add Command
	addCmd.Flags().StringVarP(&KeyName, "key-name", "n", "", "Name of the key")
	addCmd.Flags().StringVarP(&KeyValue, "key-value", "v", "", "Value of the key")
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
	keystore.LoadEncryptedYaml("keys.yaml.enc", &data, []byte("1234"))
	// keystore.LoadYaml("keys.yaml", &data)

	key, _ := cmd.Flags().GetString("key-name")
	value, _ := cmd.Flags().GetString("key-value")

	data[key] = value
	yamlData := keystore.DumpYaml("keys.yaml.enc", &data)
	keystore.EncryptKeystore("keys.yaml.enc", yamlData, []byte("1234"))
}
