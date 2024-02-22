package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/sethvargo/go-password/password"
	"gopkg.in/yaml.v2"
)

// File locations
var ConfigFile string
var KeystoreFile string

// Secret information
var PasswdEnvVar string
var Password string

// Config loaded
var ConfigLoaded bool

// Config file structure
type ConfigF struct {
	Password string `yaml:"secret"`
	Keystore string `yaml:"keystorePath"`
}

func init() {
	SetDefaults()
	ConfigLoaded = LoadConfigFile(ConfigFile)
}

func LoadConfigFile(configFile string) bool {
	// Load config file
	data := ConfigF{}

	configLoaded := false

	f, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Cannot open the configuration file '%s'!\n", configFile)
	} else {
		err = yaml.Unmarshal(f, &data)
		if err != nil {
			fmt.Printf("Cannot parse the configuration file '%s'!\n", configFile)
		} else {
			configLoaded = true
		}
	}

	// Overwrite the defaults with the values in the config file
	if configLoaded {
		Password = data.Password
		KeystoreFile = data.Keystore
	}

	// Override the password if given by environment variable
	if val, exists := os.LookupEnv(PasswdEnvVar); exists && val != "" {
		Password = val
	}

	return configLoaded
}

func SetDefaults() {
	// Set default password environment variable
	PasswdEnvVar = "JTB_TOTP_SECRET"

	// Get the default directories
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	keystoreDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Build default file locations
	ConfigFile = userConfigDir + "/jtb-totp/jtb-totp.conf"
	KeystoreFile = keystoreDir + "/.local/share/jtb-totp/keystore.enc"
}

func CreateConfigFile(pass string, keystorePath string) {
	// Create directories for keystore and config file
	configDir := path.Dir(ConfigFile)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatal(err)
	}

	keystoreDir := path.Dir(keystorePath)
	if err := os.MkdirAll(keystoreDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Set the password, generate one if not exist
	if pass == "" {
		pass, _ = password.Generate(32, 10, 0, false, false)
	}

	if keystorePath == "" {
		keystorePath = KeystoreFile
	}

	// Create config file data structure
	data := ConfigF{
		Password: pass,
		Keystore: keystorePath,
	}

	yamlData, err := yaml.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Save config file
	if err := os.WriteFile(ConfigFile, yamlData, 0400); err != nil {
		log.Fatal(err)
	}

	// Set global password for keystore creation
	Password = pass
}
