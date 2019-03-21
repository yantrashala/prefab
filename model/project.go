package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	prefab "github.com/yantrashala/prefab/common"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/yaml.v2"
)

// Project parameters
type Project struct {
	Name           string
	PID            uint32
	LocalDirectory string
	Environments   map[string]Environment
	Applications   map[string]Application
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
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
	p.LocalDirectory = path
	return nil
}

// AddEnvironment Validates and adds to teh Environment collection
func (p *Project) AddEnvironment(env Environment) error {
	//TODO: validate environment
	p.Environments[env.Name] = env

	envdir := path.Join(p.LocalDirectory, p.Name, "environments", env.Name)
	//TODO: validate path

	fmt.Println("creating env in dir: ", envdir)
	// Clones the repository into the given dir, just as a normal git clone does
	if _, err := git.PlainClone(envdir, false, &git.CloneOptions{
		URL:      env.Repo,
		Progress: os.Stdout,
	}); err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil
}

// AddApplication Validates and adds to teh Environment collection
func (p *Project) AddApplication(app Application) error {
	//TODO: validate environment
	p.Applications[app.Name] = app

	appdir := path.Join(p.LocalDirectory, p.Name, "apps", app.Name)
	//TODO: validate path

	fmt.Println("creating app in dir: ", appdir)
	// Clones the repository into the given dir, just as a normal git clone does
	if _, err := git.PlainClone(appdir, false, &git.CloneOptions{
		URL:      app.Repo,
		Progress: os.Stdout,
	}); err != nil {
		log.Fatalf("error: %v", err)
	}
	return nil
}

// SaveProject will the save the vales in the project struct
func (p *Project) SaveProject() error {
	projectFileName := path.Join(p.LocalDirectory, p.Name, "project.yaml")
	d, err := yaml.Marshal(p)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Println("Writing project to: ", projectFileName)
	fmt.Printf("--- project dump:\n%s\n\n", string(d))
	err = ioutil.WriteFile(projectFileName, d, 0644)
	return err
}

// CurrentProject values
var CurrentProject *Project

func init() {
	apps := make(map[string]Application)
	envs := make(map[string]Environment)
	conns := make(map[string]Connection)
	name := prefab.GenerateName(false)
	CurrentProject = &Project{Name: name, Environments: envs, Applications: apps, Connections: conns}
}
