package config

import (
	"fmt"
	"log"
	"os"

	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

var ConfigDir string
var ConfigFile string
var DataDir string
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

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	DataDir = userHomeDir + "/.local/share/jtb-totp"
	ConfigDir = userConfigDir + "/jtb-totp"
	ConfigFile = ConfigDir + "/jtb-totp.conf"
	KeystoreFile = DataDir + "/keystore.enc"
	PasswdEnvVar = "JTB_TOTP_SECRET"

	var val string
	var exists bool

	if val, exists = os.LookupEnv(PasswdEnvVar); !exists {
		f, err := os.ReadFile(ConfigFile)
		if err != nil {
			Password = ""
			return
		}
		var data ConfigF
		if err := yaml.Unmarshal(f, &data); err != nil {
			log.Fatal(err)
		}
		Password = data.Password
	} else {
		Password = val
	}
}

func CreateConfigFile() {
	var err error
	var pw string
	val, present := os.LookupEnv(PasswdEnvVar)
	if !present {
		fmt.Println("Not present!")
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

	Password = pw
}
