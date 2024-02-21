package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	cmd.RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize keystore and settings",
	Run:   initCommand,
}

func initCommand(cmd *cobra.Command, args []string) {
	err := os.MkdirAll(config.ConfigDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	config.CreateConfigFile()

	data := map[string]string{}
	yamlData := keystore.DumpYaml(&data)
	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Println("Initialized keystore and configuration files!")
	fmt.Printf("- Config file:   %s\n- Keystore file: %s\n", config.ConfigFile, config.KeystoreFile)
}
