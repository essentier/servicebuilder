package servicebuilder

// type serviceBuildErr struct {
// 	serviceName string
// 	err         error
// }

// type servicesBuildErr struct {
// 	errors []serviceBuildErr
// }

// func (s *servicesBuildErr) Error() string {
// 	errStr := ""
// 	for _, err := range s.errors {
// 		errStr += err.serviceName + " failed to build with error: " + err.err.Error() + "\n"
// 	}
// 	return errStr
// }

// func BuildAllInConfig(config config.Model) error {
// 	builder, err := createServicesBuilder(config)
// 	if err != nil {
// 		return err
// 	}
// 	return builder.buildAllServices()
// }

// func createServicesBuilder(configModel config.Model) (*servicesBuilder, error) {
// 	// configModel, err := config.GetConfig()
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	sb := &servicesBuilder{config: configModel}
// 	err := sb.init()
// 	return sb, err
// }

// type servicesBuilder struct {
// 	config      config.Model
// 	token       string
// 	controllers buildControllers
// }

// func (p *servicesBuilder) init() error {
// 	cloudProvider := p.config.CloudProvider
// 	token, err := model.LoginToEssentier(cloudProvider.Url, cloudProvider.Username, cloudProvider.Password)
// 	p.token = token
// 	if err != nil {
// 		return err
// 	}

// 	controllers, err := createBuildControllers(p.config)
// 	p.controllers = controllers
// 	return err
// }

// // By the time this method returns, the service will have been built.
// func (p *servicesBuilder) ensureServiceIsBuilt(serviceName string) error {
// 	return p.controllers.ensureServiceIsBuilt(serviceName, p.config.CloudProvider.Url, p.token)
// }

// func (p *servicesBuilder) buildAllServices() error {

// 	errs := p.buildServices(allServices)
// 	if len(errs) == 0 {
// 		return nil
// 	} else {
// 		return &servicesBuildErr{errors: errs}
// 	}
// }

// func (p *servicesBuilder) buildServices(allServices map[string]config.Service) []serviceBuildErr {
// 	resultsChan := make(chan serviceBuildErr)
// 	for _, serviceConfig := range allServices { //build services concurrently
// 		go buildService(serviceConfig, p.config.CloudProvider.Url, p.token, resultsChan)
// 	}

// 	failedBuilds := []serviceBuildErr{}
// 	for i := 0; i < len(allServices); i++ {
// 		r := <-resultsChan
// 		if r.err != nil {
// 			failedBuilds = append(failedBuilds, r)
// 		}
// 	}
// 	return failedBuilds
// }

// func buildService(serviceConfig config.Service, providerUrl string,
// 	token string) {
// 	serviceBuilder := createServiceBuilder(serviceConfig, providerUrl, token)
// 	err := serviceBuilder.buildService()
// 	resultsChan <- serviceBuildErr{serviceName: serviceConfig.ServiceName, err: err}
// }

////////////
// func collectSourceServices(configModel config.Model, serviceMap map[string]config.Service) error {
// 	for serviceName, serviceConfig := range configModel.Services {
// 		if !serviceConfig.IsSourceProject() {
// 			//log.Printf("Service %v is not a source project. Skip.", serviceName)
// 			continue
// 		}

// 		if _, exists := serviceMap[serviceConfig.ServiceName]; exists {
// 			//log.Printf("Service %v is already visited. Skip.", serviceName)
// 			continue // Service already visited. Skip.
// 		}

// 		log.Printf("Found new source service %v.", serviceName)
// 		serviceMap[serviceName] = serviceConfig

// 		//The service is a source project. It may have its own spickspan config.
// 		fullFileName := filepath.Join(serviceConfig.ProjectSrcRoot, config.SpickSpanConfigFile)
// 		//log.Printf("Check if service %v has spickspan file %v.", serviceName, fullFileName)
// 		_, err := os.Stat(fullFileName)
// 		if os.IsNotExist(err) {
// 			// The service does not have its own spickspan conifg. Move on.
// 			//log.Printf("Service %v does not have its own spickspan config.", serviceName)
// 			continue
// 		}

// 		newConfigModel, err := config.ParseConfigFile(fullFileName)
// 		if err != nil {
// 			return err
// 		}

// 		err = collectSourceServices(newConfigModel, serviceMap)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
