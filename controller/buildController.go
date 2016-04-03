package controller

import (
	"log"

	"github.com/essentier/servicebuilder/builder"
	"github.com/go-errors/errors"
)

type BuildController interface {
	EnsureServiceBuilt(serviceName string) error
}

// func BuildAllInConfig(config config.Model) error {
// 	return servicebuilder.BuildAllInConfig(config)
// }

// A build controller controls the building of a set of source services.
// It uses a service builder to build services.
// It makes sure that every source service is built at most once.
// It is safe for concurrent access.
type buildController struct {
	builder      builder.ServiceBuilder
	buildDataMap map[string]buildData // The set of services under control.
}

func (c buildController) EnsureServiceBuilt(serviceName string) error {
	buildData, exists := c.buildDataMap[serviceName]
	if !exists {
		return errors.Errorf("Could not find build data of %s", serviceName)
	}

	var err error
	buildData.once.Do(func() {
		log.Printf("Will build service")
		sourceProject, err := buildData.createSourceProject()
		if err != nil {
			return
		}

		err = c.builder.BuildService(serviceName, sourceProject)
	})
	return err
}

// func createBuildControllers() buildControllers {
// 	return buildControllers{controllerMap: map[string]buildController{}}
// }

// var dummyOnce sync.Once

// func (bcs *buildControllers) getServiceOnce(serviceName string) (*sync.Once, error) {
// 	bc, exists := bcs.controllerMap[serviceName]
// 	if exists {
// 		return bc.once, nil
// 	} else {
// 		return &dummyOnce, nil
// 	}
// }

// func (sb *servicesBuild) contains(serviceName string) bool {
// 	_, exists := sb.buildMap[serviceName]
// 	return exists
// }
