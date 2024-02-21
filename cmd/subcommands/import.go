package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"jamesthebard/jtb-totp/cmd"
	"jamesthebard/jtb-totp/keystore"
)

func init() {
	var Overwrite bool

	importCmd.Flags().BoolVarP(&Overwrite, "overwrite", "o", false, "overwrite the keys in current keystore")

	cmd.RootCmd.AddCommand(importCmd)
}

var importCmd = &cobra.Command{
	Use:   "import [file to import]",
	Short: "Import keystore from YAML/JSON file",
	Args:  cobra.ExactArgs(1),
	Run:   importCommand,
}

func importCommand(cmd *cobra.Command, args []string) {
	overwrite, _ := cmd.Flags().GetBool("overwrite")
	keystore.ImportKeystore(args[0], overwrite)
	fmt.Printf("Updated keystore with imported keys.\n")
}
