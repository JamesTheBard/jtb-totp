package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize keystore and settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}
