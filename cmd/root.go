package cmd

import (
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
}

func initConfig() {}

func rootCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
