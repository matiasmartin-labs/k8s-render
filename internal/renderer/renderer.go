package renderer

import (
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/matiasmartin-labs/k8s-render/internal/config"
	"github.com/matiasmartin-labs/k8s-render/internal/utils"
)

func RenderK8sManifests(outputDir string, platformConfig *config.PlatformConfig, vars utils.VarsFlag) error {
	logger := utils.GetLogger()

	pattern := filepath.Join("./templates", "*.template.yml")

	logger.Infof("Searching for template files in: %s", pattern)

	files, err := filepath.Glob(pattern)
	if err != nil {
		logger.Errorf("Error finding template files: %v", err)
		return err
	}

	logger.Infof("Found %d template files", len(files))
	
	if len(files) == 0 {
		logger.Warn("No template files found. Exiting.")
		return nil
	}

	logger.Infof("Rendering templates to: %s", outputDir)

	for _, file := range files {
		logger.Debugf("Processing template file: %s", file)

		tmpl, err := template.ParseFiles(file)
		if err != nil {
			logger.Errorf("Error parsing template file %s: %v", file, err)
			return err
		}

		outputFileName := filepath.Base(file)
		outputFileName = outputFileName[:len(outputFileName)-len(".template.yml")] + ".yml"

		outputFilePath := filepath.Join(outputDir, outputFileName)

		logger.Debugf("Rendering template to: %s", outputFilePath)

		if err := os.MkdirAll(filepath.Dir(outputFilePath), os.ModePerm); err != nil {
			logger.Errorf("Error creating output directory: %v", err)
			return err
		}

		outputFile, err := os.Create(outputFilePath)
		if err != nil {
			logger.Errorf("Error creating output file %s: %v", outputFilePath, err)
			return err
		}

		defer func() {
			if err := outputFile.Close(); err != nil {
				logger.Errorf("Error closing output file %s: %v", outputFilePath, err)
			}
		}()

		data := utils.StructToMap(platformConfig)
		data["Vars"] = vars
		data["RenderTime"] = time.Now().Format(time.RFC3339)

		if err := tmpl.Execute(outputFile, data); err != nil {
			logger.Errorf("Error executing template %s: %v", file, err)
			return err
		}
	}

	logger.Infof("Rendering completed successfully. Manifests written to %s", outputDir)
	return nil

}
