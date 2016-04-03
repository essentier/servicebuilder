package builder

import (
	"log"
	"strings"

	"github.com/essentier/gopencils"
)

type nativeBuilder interface {
	buildService(serviceName string, token string) error
	url() string
}

func createDefaultNativeBuilder(providerUrl string) nativeBuilder {
	if strings.HasSuffix(providerUrl, "/") {
		providerUrl = providerUrl[:len(providerUrl)-1]
	}
	return nomockBuilder{nomockApi: gopencils.Api(providerUrl)}
}

type nomockBuilder struct {
	nomockApi *gopencils.Resource
}

func (p nomockBuilder) url() string {
	return p.nomockApi.Api.BaseUrl.String() + "/nomockbuilder"
}

func (p nomockBuilder) buildService(serviceName string, token string) error {
	log.Printf("building service %v on nomock builder", serviceName)
	var result interface{}
	builderResource := p.nomockApi.NewChildResource("nomockbuilder/build/"+serviceName, &result)
	builderResource.SetHeader("Authorization", "Bearer "+token)
	_, err := builderResource.Get()
	return err
}
