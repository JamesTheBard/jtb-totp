package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/config"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var force bool
	var keystorePath string
	var password string

	initCmd.Flags().BoolVarP(&force, "force", "f", false, "force re-initialization of config and keystore")
	initCmd.Flags().StringVarP(&keystorePath, "keystore", "k", "", "location of new keystore path")
	initCmd.Flags().StringVarP(&password, "password", "p", "", "encrypt datastore with user-defined password")

	cmd.RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize keystore and settings",
	Run:   initCommand,
}

func initCommand(cmd *cobra.Command, args []string) {
	// Determine if the program has already been initialized and if it has, stop
	// the program from re-initializing unless the 'force' option is given.
	config.SetDefaults()
	force, _ := cmd.Flags().GetBool("force")
	if configLoaded, _ := config.LoadConfigFile(config.ConfigFile); configLoaded && !force {
		fmt.Println("The program has already been configured, you must use the --force option to re-initialize.")
		os.Exit(1)
	}

	password, _ := cmd.Flags().GetString("password")
	keystorePath, _ := cmd.Flags().GetString("keystore")

	// Where those directories should be created
	config.CreateConfigFile(password, keystorePath)

	// Create a new empty encrypted keystore
	data := map[string]string{}
	yamlData := keystore.DumpYaml(&data)
	keystore.EncryptKeystore(config.KeystoreFile, yamlData, []byte(config.Password))
	fmt.Println("Initialized keystore and configuration files!")
	fmt.Printf("- Config file:   %s\n- Keystore file: %s\n", config.ConfigFile, config.KeystoreFile)
}
