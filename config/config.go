package config

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Arl          string `toml:"arl"`
	LicenseToken string `toml:"license_token"`
	DestDir      string `toml:"dest_dir"`
	Iv           string `toml:"iv"`
	PreKey       string `toml:"pre_key"`
}

func GetConfig() (Configuration, error) {
	var err error
	var config Configuration

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if len(configDir) == 0 {
		homedir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		configDir = homedir + "/.config/"
	}
	configPath := configDir + "/deezer-flac-download/config.toml"

	_, err = toml.DecodeFile(configPath, &config)
	if err != nil {
		return Configuration{}, err
	}
	if len(config.Arl) == 0 {
		return Configuration{}, errors.New("please provide a value for the 'arl' field in the config file")
	}
	if len(config.LicenseToken) == 0 {
		return Configuration{}, errors.New("please provide a value for the 'license_token' field in the config file")
	}
	if len(config.DestDir) == 0 {
		return Configuration{}, errors.New("please provide a value for the 'dest_dir' field in the config file")
	}
	if len(config.PreKey) == 0 {
		return Configuration{}, errors.New("please provide a value for the 'pre_key' field in the config file")
	}
	if len(config.Iv) == 0 {
		return Configuration{}, errors.New("please provide a value for the 'iv' field in the config file")
	}
	return config, nil
}
