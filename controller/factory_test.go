package controller

import (
	"testing"

	"github.com/essentier/spickspan/config"
	"github.com/stretchr/testify/assert"
)

func TestCollectBuildData(t *testing.T) {
	assert := assert.New(t)
	configModel := config.CreateTestConfigModel()
	buildDataMap := map[string]buildData{}

	err := collectBuildData(configModel, buildDataMap)
	assert.Nil(err)

	_, exists := buildDataMap["todo-rest"]
	assert.True(exists)

	_, exists = buildDataMap["mongodb"]
	assert.False(exists)
}
