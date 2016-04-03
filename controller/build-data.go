package controller

import (
	"path/filepath"
	"sync"

	"github.com/essentier/nomockutil"
	"github.com/essentier/servicebuilder/scm"
	"github.com/essentier/servicebuilder/scm/git"
	"github.com/essentier/spickspan/config"
	"github.com/go-errors/errors"
)

type buildData struct {
	serviceConfig config.Service
	once          *sync.Once
}

func (d buildData) createSourceProject() (scm.Project, error) {
	projectDir := d.serviceConfig.ProjectSrcRoot
	gitDir := filepath.Join(projectDir, ".git")
	if !nomockutil.Exists(gitDir) {
		return nil, errors.Errorf("Project %v is not initialized with git. Use 'git init' to initialize the project.", projectDir)
	}

	return git.CreateDefaultGitProject(projectDir), nil
}
