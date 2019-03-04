package model

import (
	"errors"
	"io/ioutil"

	prefab "github.com/yantrashala/prefab/common"
)

// Project parameters
type Project struct {
	Name           string
	PID            uint32
	LocalDirectory string
	Environments   map[string]Environment
	Connections    map[string]Connection
}

// SetProjectName Validates and Sets the projectName
func (p *Project) SetProjectName(name string) error {
	if len(name) < 3 {
		return errors.New("Invalid project name, has to atleast 3 characters long")
	}
	p.Name = name
	return nil
}

// SetLocalDirectory Validates and Sets the local path where the generated files will be placed.
func (p *Project) SetLocalDirectory(path string) error {
	if _, err := ioutil.ReadDir(path); err != nil {
		return errors.New("Invalid path for localdirectory, enable to read(check permissions)")
	}
	p.LocalDirectory = path
	return nil
}

// AddEnvironment Validates and adds to teh Environment collection
func (p *Project) AddEnvironment(env Environment) error {
	//TODO: validate environment
	p.Environments[env.Name] = env
	return nil
}

// CurrentProject values
var CurrentProject *Project

func init() {
	envs := make(map[string]Environment)
	conns := make(map[string]Connection)
	name := prefab.GenerateName(false)
	CurrentProject = &Project{Name: name, Environments: envs, Connections: conns}
}
