package cmd

import (
	"fmt"
	"os"
	"sort"

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
	err := keystore.LoadEncryptedYaml(config.KeystoreFile, &data, []byte(config.Password))
	if err != nil {
		fmt.Printf("Could not open/process the keystore: %s\n", err)
		os.Exit(1)
	}

	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k)
	}
}
