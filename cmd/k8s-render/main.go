package main

import (
	"flag"

	"github.com/matiasmartin-labs/k8s-render/internal/config"
	"github.com/matiasmartin-labs/k8s-render/internal/renderer"
	"github.com/matiasmartin-labs/k8s-render/internal/utils"
)

var (
	inputPath  string
	outputPath string
	logLevel   string
	vars       = utils.VarsFlag{}
)

func init() {
	flag.StringVar(&inputPath, "input", "./k8s", "Path to the input file")
	flag.StringVar(&outputPath, "output", "./k8s/manifests", "Path to the output file")
	flag.StringVar(&logLevel, "log-level", "info", "Log level (debug, info, warn, error, fatal, panic)")
	flag.Var(&vars, "var", "Variables to pass to the templates in key=value format")
	flag.Parse()
}

func main() {

	logger := utils.NewLogger(logLevel)

	logger.Info("Starting k8s-render...")
	logger.Debugf("Input path: %s", inputPath)
	logger.Debugf("Output path: %s", outputPath)
	logger.Debugf("Log level: %s", logLevel)
	logger.Debugf("Variables: %v", vars)

	logger.Info("Rendering Kubernetes manifests...")

	// Here you would add the logic to read the input file, process it, and write the output file.
	platformConfig, err := config.LoadPlatformConfig(inputPath)
	if err != nil {
		return
	}

	// For now, we'll just simulate the rendering process with a sleep.
	err = renderer.RenderK8sManifests(outputPath, platformConfig, vars)
	if err != nil {
		return
	}

	logger.Info("Rendering completed successfully.")
	logger.Infof("Manifests written to %s", outputPath)
	logger.Info("Exiting k8s-render.")
	logger.Info("Goodbye!")
}
