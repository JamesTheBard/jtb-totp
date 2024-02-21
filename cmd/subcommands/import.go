package cmd

import (
	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
)

func init() {
	cmd.RootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import keystore from YAML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}
