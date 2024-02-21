package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	cmd.RootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use:   "export [file to save export to]",
	Short: "Export keystore to YAML file or standard out",
	Run:   exportCommand,
}

func exportCommand(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		keystore.ExportKeystore("", true)
	} else {
		keystore.ExportKeystore(args[0], false)
		fmt.Printf("Keystore exported to '%s'.\n", args[0])
	}
}
