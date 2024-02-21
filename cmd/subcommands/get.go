package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Generate TOTP code from key",
	Run:   func(cmd *cobra.Command, args []string) {},
}
