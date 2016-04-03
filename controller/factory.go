package controller

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/essentier/servicebuilder/builder"
	"github.com/essentier/spickspan/config"
)

func CreateDefaultBuildController(configModel config.Model) (BuildController, error) {
	serviceBuilder, err := builder.CreateDefaultServiceBuilder(configModel.CloudProvider)
	if err != nil {
		return buildController{}, err
	}
	return CreateBuildController(configModel, serviceBuilder)
}

func CreateBuildController(configModel config.Model, serviceBuilder builder.ServiceBuilder) (BuildController, error) {
	buildDataMap := map[string]buildData{}
	err := collectBuildData(configModel, buildDataMap)
	return buildController{builder: serviceBuilder, buildDataMap: buildDataMap}, err
}

func collectBuildData(configModel config.Model, buildDataMap map[string]buildData) error {
	for serviceName, serviceConfig := range configModel.Services {
		if !serviceConfig.IsSourceProject() {
			continue
		}

		log.Printf("Found new source service %v.", serviceName)
		if _, exists := buildDataMap[serviceName]; exists {
			continue
		}

		var newOnce sync.Once
		buildDataMap[serviceName] = buildData{serviceConfig: serviceConfig, once: &newOnce}

		//The service is a source project. It may have its own spickspan config.
		fullFileName := filepath.Join(serviceConfig.ProjectSrcRoot, config.SpickSpanConfigFile)
		//log.Printf("Check if service %v has spickspan file %v.", serviceName, fullFileName)
		_, err := os.Stat(fullFileName)
		if os.IsNotExist(err) {
			// The service does not have its own spickspan conifg. Move on.
			//log.Printf("Service %v does not have its own spickspan config.", serviceName)
			continue
		}

		newConfigModel, err := config.ParseConfigFile(fullFileName)
		if err != nil {
			return err
		}

		err = collectBuildData(newConfigModel, buildDataMap)
		if err != nil {
			return err
		}
	}
	return nil
}
