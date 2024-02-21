package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var ConfigDir string
var ConfigFile string
var KeystoreFile string
var Password string

type ConfigF struct {
	Password string `yaml:"password"`
}

func init() {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	ConfigDir = userConfigDir + "/jtb-totp"
	ConfigFile = ConfigDir + "/jtb-totp.conf"
	KeystoreFile = ConfigDir + "/keystore.enc"

	f, err := os.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}

	var data ConfigF
	if err := yaml.Unmarshal(f, &data); err != nil {
		log.Fatal(err)
	}
	Password = data.Password
}
