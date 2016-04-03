package git_test

import (
	"testing"

	"github.com/essentier/servicebuilder/scm/git"
	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	projectDir := "/src/github.com/example"
	recordCmd := createRecordCmd()
	cmdRunner := git.CreateGitCmdRunner(projectDir, recordCmd)
	cmdRunner.RunCmd("stash")
	expected := "git -C " + projectDir + " stash"
	assert.Equal(t, expected, recordCmd.lastCommand)

	cmdRunner.RunCmd("pull", "-s", "ours", "remoteUrl", "branchName")
	expected = "git -C " + projectDir + " pull -s ours remoteUrl branchName"
	assert.Equal(t, expected, recordCmd.lastCommand)
}
