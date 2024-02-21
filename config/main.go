package config

import (
	"log"
	"os"

	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

var ConfigDir string
var ConfigFile string
var KeystoreFile string
var Password string
var PasswdEnvVar string

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
	PasswdEnvVar = "JTB_TOTP_SECRET"

	f, err := os.ReadFile(ConfigFile)
	if err != nil {
		Password = os.Getenv(PasswdEnvVar)
	} else {
		var data ConfigF
		if err := yaml.Unmarshal(f, &data); err != nil {
			log.Fatal(err)
		}
		Password = data.Password
	}
}

func CreateConfigFile() {
	var err error
	var pw string
	val, present := os.LookupEnv(PasswdEnvVar)
	if !present {
		pw, err = password.Generate(32, 10, 0, false, false)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		pw = val
	}

	data := ConfigF{
		Password: pw,
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(ConfigFile, yamlData, 0400); err != nil {
		log.Fatal(err)
	}
}
