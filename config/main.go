package config

import (
	"log"
	"os"

	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

var ConfigFile string
var DataDir string
var KeystoreFile string
var Password string
var PasswdEnvVar string

type ConfigF struct {
	Password string `yaml:"secret"`
	Keystore string `yaml:"keystorePath,omitempty"`
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

	KeystoreFile = userHomeDir + "/.local/share/jtb-totp/keystore.enc"
	ConfigFile = userConfigDir + "/jtb-totp/jtb-totp.conf"
	PasswdEnvVar = "JTB_TOTP_SECRET"

	f, err := os.ReadFile(ConfigFile)
	if err == nil {
		data := ConfigF{}
		yaml.Unmarshal(f, &data)

		Password = data.Password
		KeystoreFile = data.Keystore
	}

	if val, exists := os.LookupEnv(PasswdEnvVar); exists {
		Password = val
	}
}

func CreateConfigFile(pass string, keystorePath string) {
	if len(pass) == 0 {
		pass, _ = password.Generate(32, 10, 0, false, false)
	}

	if len(keystorePath) == 0 {
		keystorePath = KeystoreFile
	}

	data := ConfigF{
		Password: pass,
		Keystore: keystorePath,
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(ConfigFile, yamlData, 0400); err != nil {
		log.Fatal(err)
	}

	Password = pass
}
