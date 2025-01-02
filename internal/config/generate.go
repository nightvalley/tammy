// toml

package config

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

func (config *Config) generateConfig() {}
