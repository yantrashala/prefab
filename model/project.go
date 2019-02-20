package model

import (
	"errors"
	"io/ioutil"
)

// Project parameters
type Project struct {
	Name           string
	PID            uint32
	LocalDirectory string
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

// CurrentProject values
var CurrentProject *Project

func init() {
	CurrentProject = &Project{}
}
