package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove key from keystore",
	Run:   func(cmd *cobra.Command, args []string) {},
}
