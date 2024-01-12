package rating

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config is a structure containing data from the config.yaml file
type Config struct {
	Lichess struct {
		URL         string `yaml:"url"`
		DefaultUser string `yaml:"defaultUser"`
	} `yaml:"lichess"`
	USCF struct {
		URL          string `yaml:"url"`
		DefaultUser  string `yaml:"defaultUser"`
		DefaultState string `yaml:"defaultState"`
	} `yaml:"USCF"`
}

var (
	DEFAULT_DATA_GETTER = ReadConfigFile
	// Override this for unit tests
	DATA_GETTER = DEFAULT_DATA_GETTER
)

// LoadsConfig reads the configuration YAML file and returns a
// configuration object.
//
// The configuration data can be mocked for unit testing by
// setting DATA_GETTER to a function that returns ([]byte, error)
func LoadConfig() (*Config, error) {
	var config Config
	body, err := DATA_GETTER()
	defer func() {
		DATA_GETTER = DEFAULT_DATA_GETTER
	}()
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

// ReadConfigFile reads $HOME/.config/rating/config.yaml and returns
// its contents
func ReadConfigFile() ([]byte, error) {
	configDir, _ := os.UserConfigDir()
	filename := filepath.Join(configDir, "chess-rating", "config.yaml")
	body, err := os.ReadFile(filename)
	return body, err
}
