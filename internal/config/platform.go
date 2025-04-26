package config

import (
	"fmt"
	"os"

	"github.com/matiasmartin-labs/k8s-render/internal/utils"
	"gopkg.in/yaml.v3"
)

const platformFileName = "platform.yaml"

type Application struct {
	Name   string `yaml:"name"`
	Port   int    `yaml:"port"`
	PartOf string `yaml:"part-of"`
}

type Network struct {
	Host string `yaml:"host"`
}

type PlatformConfig struct {
	App       Application `yaml:"app"`
	Namespace string      `yaml:"namespace"`
	Network   Network     `yaml:"network"`
}

func LoadPlatformConfig(path string) (*PlatformConfig, error) {
	logger := utils.GetLogger()
	logger.Info("Loading platform configuration...")

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", path, platformFileName))
	if err != nil {
		logger.Errorf("Error reading platform configuration file: %v", err)
		return nil, err
	}

	var config PlatformConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		logger.Errorf("Error unmarshalling platform configuration: %v", err)
		return nil, err
	}

	logger.Debugf("Loaded platform configuration: %+v", config)

	return &config, nil
}
