package config

import (
	"gopkg.in/yaml.v3"
	"iot-messaging/cmd/api/model"
	"os"
)

const Path = "./resources/application.yml"

func LoadConfig() (model.MessageConfig, error) {
	data, err := os.ReadFile(Path)
	if err != nil {
		return model.MessageConfig{}, err
	}
	var config model.MessageConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return model.MessageConfig{}, err
	}
	return config, nil
}
