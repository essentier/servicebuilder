package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var urls = []string{
	"https://nomock.essentier.com",
	"https://nomock.essentier.com/"}

func TestGetServiceRepoUrl(t *testing.T) {
	for _, providerUrl := range urls {
		token := "jlasdlk"
		sb := serviceBuilder{nativeBuilder: createDefaultNativeBuilder(providerUrl), token: token}
		repoUrl := sb.getServiceDepoUrl("todo-app")
		assert.Equal(t, "https://"+token+":@nomock.essentier.com/nomockbuilder/todo-app", repoUrl)
	}
}
