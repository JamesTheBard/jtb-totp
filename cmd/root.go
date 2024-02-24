package cmd

import (
	"fmt"
	"jamesthebard/jtb-totp/config"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "jtb-totp",
	Short: "JTB TOTP is makes cool numbers",
	Long:  `JTB-TOTP is a quick-and-dirty program that generates TOTP tokens and manages TOTP keys via the command-line.`,
	Run:   rootCommand,
}

func Execute() error {
	return RootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.CompletionOptions.HiddenDefaultCmd = true

	var versionFlag bool

	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "display the version and exit")
}

func initConfig() {}

func rootCommand(cmd *cobra.Command, args []string) {
	if isVersion, _ := cmd.Flags().GetBool("version"); isVersion {
		fmt.Printf("jtb-totp %s\n", config.Version)
		os.Exit(0)
	}

	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
