package controller_test

import (
	"testing"
	"time"

	"github.com/essentier/servicebuilder/controller"
	"github.com/essentier/servicebuilder/scm"
	"github.com/essentier/spickspan/config"
	"github.com/stretchr/testify/mock"
)

type mockServiceBuilder struct {
	mock.Mock
}

func (m *mockServiceBuilder) BuildService(serviceName string, project scm.Project) error {
	m.Called(serviceName, project)
	return nil
	//return args.Error(0)
}

func TestEnsureServiceIsBuilt(t *testing.T) {
	mockBuilder := new(mockServiceBuilder)
	configModel := config.CreateTestConfigModel()
	bc, _ := controller.CreateBuildController(configModel, mockBuilder)

	go bc.EnsureServiceBuilt("todo-rest")
	go bc.EnsureServiceBuilt("todo-rest")
	go bc.EnsureServiceBuilt("todo-rest")

	time.Sleep(1000 * time.Millisecond)
	// assert that the expectations were met
	mockBuilder.AssertNumberOfCalls(t, "BuildService", 1)
}
