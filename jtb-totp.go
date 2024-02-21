package main

import (
	"jamesthebard/jtb-totp/cmd"
	_ "jamesthebard/jtb-totp/cmd/subcommands"
)

type TotpKey struct {
	Name string `yaml:"name"`
	Key  string `yaml:"key"`
}

type TotpKeyStore struct {
	KeyStore []TotpKey
}

func main() {
	cmd.Execute()
}
