// toml

package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DefaultFileSizeFlag bool
	DefaultHiddenFiles  bool
	DefaultForm         string

	DefaultTreePreset   string
	TreeEnumeratorType  string
	TreeEnumeratorColor int

	DefaultListPreset   string
	ListEnumeratorType  string
	ListEnumeratorColor int

	DefaultTablePreset string
	TableBorder        string
	TableBorderColor   int
}

func (config *Config) generateConfig() {
	UserHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Panicf("Failed to get user home directory: %v", err)
	}
	viper.SetConfigFile(UserHomeDir + ".config/counlines.toml")
}
