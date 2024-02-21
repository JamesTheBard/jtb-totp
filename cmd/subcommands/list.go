package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all keys by name in keystore",
	Run:   listCommand,
}

func listCommand(cmd *cobra.Command, args []string) {
	data := make(map[string]string)
	keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))

	for k := range data {
		fmt.Println(k)
	}
}
