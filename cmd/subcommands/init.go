package cmd

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var password string
	var keystorePath string
	var force bool

	initCmd.Flags().StringVarP(&password, "password", "p", "", "encrypt datastore with user-defined password")
	initCmd.Flags().StringVarP(&keystorePath, "keystore", "k", "", "location of new keystore path")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-initialization of config and keystore")

	cmd.RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize keystore and settings",
	Run:   initCommand,
}

func initCommand(cmd *cobra.Command, args []string) {
	if force, _ := cmd.Flags().GetBool("force"); !force {
		if _, err := os.Stat(config.ConfigFile); err == nil {
			fmt.Println("The program has already been initialized, you must use the --force option to re-initialize.")
			os.Exit(1)
		}

		keystoreVal, _ := cmd.Flags().GetString("keystore")
		if len(keystoreVal) == 0 {
			keystoreVal = config.KeystoreFile
		}

		if _, err := os.Stat(keystoreVal); err == nil {
			fmt.Println("The program has already been initialized, you must use the --force option to re-initialize.")
			os.Exit(1)
		}
	}

	configDir := path.Dir(config.ConfigFile)
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	keystoreDir := path.Dir(config.KeystoreFile)
	err = os.MkdirAll(keystoreDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	password, _ := cmd.Flags().GetString("password")
	keystorePath, _ := cmd.Flags().GetString("keystore")
	config.CreateConfigFile(password, keystorePath)

	data := map[string]string{}
	yamlData := keystore.DumpYaml(&data)
	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Println("Initialized keystore and configuration files!")
	fmt.Printf("- Config file:   %s\n- Keystore file: %s\n", config.ConfigFile, config.KeystoreFile)
}
