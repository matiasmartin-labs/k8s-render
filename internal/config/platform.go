package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/matiasmartin-labs/k8s-render/internal/utils"
	"gopkg.in/yaml.v3"
)

const platformFileName = "platform.yaml"

type HealthCheckProbe struct {
	Path                string `yaml:"path" validate:"required"`
	InitialDelaySeconds int    `yaml:"initial-delay-seconds" validate:"required"`
}

type HealthCheck struct {
	Liveness  HealthCheckProbe `yaml:"liveness" validate:"required"`
	Readiness HealthCheckProbe `yaml:"readiness" validate:"required"`
}

type ResourceRequirements struct {
	CPU    string `yaml:"cpu" validate:"required"`
	Memory string `yaml:"memory" validate:"required"`
}

type Resources struct {
	Requests ResourceRequirements `yaml:"requests" validate:"required"`
	Limits   ResourceRequirements `yaml:"limits" validate:"required"`
}

type Metrics struct {
	Path string `yaml:"path" validate:"required"`
}

type Application struct {
	Name        string      `yaml:"name" validate:"required"`
	Port        int         `yaml:"port" validate:"required"`
	PartOf      string      `yaml:"part-of" validate:"required"`
	Replicas    int         `yaml:"replicas" validate:"required"`
	HealthCheck HealthCheck `yaml:"health-check" validate:"required"`
	Resources   Resources   `yaml:"resources" validate:"required"`
	Metrics     Metrics     `yaml:"metrics" validate:"required"`
}

type Network struct {
	Host string `yaml:"host" validate:"required"`
}

type Environment struct {
	Secrets []string `yaml:"secrets"`
}

type Mount struct {
	Name     string `yaml:"name" validate:"required"`
	Path     string `yaml:"mount-path" validate:"required"`
	ReadOnly bool   `yaml:"read-only" validate:"required"`
}

type SecretVolume struct {
	Name       string `yaml:"name" validate:"required"`
	Mode       int    `yaml:"default-mode" validate:"required"`
	SecretName string `yaml:"secret-name" validate:"required"`
}

type PlatformConfig struct {
	App           Application    `yaml:"app" validate:"required"`
	Namespace     string         `yaml:"namespace" validate:"required"`
	Network       Network        `yaml:"network" validate:"required"`
	Env           Environment    `yaml:"env"`
	Mounts        []Mount        `yaml:"mounts"`
	SecretVolumes []SecretVolume `yaml:"secret-volumes"`
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

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		logger.Errorf("Validation error in platform configuration: %v", err)
		return nil, err
	}

	logger.Infof("Platform configuration loaded successfully.")

	logger.Debugf("Loaded platform configuration: %+v", config)

	return &config, nil
}
