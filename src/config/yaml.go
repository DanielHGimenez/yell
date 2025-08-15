package config

import (
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	SoundPath string `yaml:"sound_path"`
}

func LoadConfig() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err := os.WriteFile(configFilePath, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() (string, error) {
	execDir, err := getExecutablePath()
	if err != nil {
		return "", err
	}
	return path.Join(execDir, "config.yaml"), nil
}
