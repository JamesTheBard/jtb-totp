package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var Overwrite bool

	importCmd.Flags().BoolVarP(&Overwrite, "overwrite", "o", false, "overwrite the current keystore (required)")
	importCmd.MarkFlagRequired("overwrite")

	cmd.RootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import [file to import]",
	Short: "Import keystore from YAML/JSON file",
	Args:  cobra.ExactArgs(1),
	Run:   importCommand,
}

func importCommand(cmd *cobra.Command, args []string) {
	keystore.ImportKeystore(args[0])
	fmt.Printf("Updated keystore with imported keys.\n")
}
