package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all names in keystore",
	Run:   func(cmd *cobra.Command, args []string) {},
}
