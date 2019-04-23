package model

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	prefab "github.com/yantrashala/prefab/common"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/filesystem"
	"gopkg.in/yaml.v2"
)

// Project parameters
type Project struct {
	Name           string
	PID            uint32
	LocalDirectory string
	Environments   map[string]Environment
	Applications   map[string]Application
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
	p.LocalDirectory = path
	return nil
}

// AddEnvironment Validates and adds to teh Environment collection
func (p *Project) AddEnvironment(env Environment) error {
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

	env.LocalDirectory = envdir
	//TODO: validate environment
	p.Environments[env.Name] = env

	os.RemoveAll(path.Join(envdir, ".git"))

	fs := osfs.New(path.Join(envdir, ".git"))
	fss := filesystem.NewStorage(fs, nil)
	git.Init(fss, fs)

	return nil
}

// AddApplication Validates and adds to teh Environment collection
func (p *Project) AddApplication(app Application) error {
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

	app.LocalDirectory = appdir
	//TODO: validate environment
	p.Applications[app.Name] = app

	os.RemoveAll(path.Join(appdir, ".git"))

	fs := osfs.New(path.Join(appdir, ".git"))
	fss := filesystem.NewStorage(fs, nil)
	git.Init(fss, fs)
	return nil
}

type tdata struct {
	Config map[string]string
	Project
}

// ApplyValues will subsitute any go templates
func (p *Project) ApplyValues() {
	// projectPath := path.Join(p.LocalDirectory, p.Name)
	for env := range p.Environments {
		envDir := p.Environments[env].LocalDirectory
		err := filepath.Walk(envDir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				// fmt.Println(path, info.Size())
				if info.IsDir() {
					return nil
				}

				i := strings.LastIndex(strings.ToLower(path), ".ptmpl")

				if i == -1 {
					return nil
				}

				if contents, ferr := ioutil.ReadFile(path); ferr == nil {
					// Create a new template and parse the letter into it.
					t := template.Must(template.New("temp").Parse(string(contents)))

					if out, oerr := os.Create(path[:i]); oerr == nil {
						data := tdata{Config: p.Environments[env].Config, Project: *p}
						t.Execute(out, data)
						defer out.Close()
					}
				} else {
					fmt.Println(ferr)
				}

				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}

	for app := range p.Applications {
		appDir := p.Applications[app].LocalDirectory
		err := filepath.Walk(appDir,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				//fmt.Println(path, info.Size())
				if info.IsDir() {
					return nil
				}

				i := strings.LastIndex(strings.ToLower(path), ".ptmpl")

				if i == -1 {
					return nil
				}

				if contents, ferr := ioutil.ReadFile(path); ferr == nil {
					// Create a new template and parse the letter into it.
					t := template.Must(template.New("temp").Parse(string(contents)))

					if out, oerr := os.Create(path[:i]); oerr == nil {
						data := tdata{Config: p.Applications[app].Config, Project: *p}
						t.Execute(out, data)
						defer out.Close()
					}
				} else {
					fmt.Println(ferr)
				}
				return nil
			})
		if err != nil {
			log.Println(err)
		}
	}
}

// GetProjectFilename gets the full path of the project settings file
func (p *Project) GetProjectFilename() string {
	return path.Join(p.LocalDirectory, p.Name, "project.yaml")
}

// SaveProject will the save the vales in the project struct
func (p *Project) SaveProject() error {
	ppath := path.Join(p.LocalDirectory, p.Name)
	if _, err := os.Stat(ppath); os.IsNotExist(err) {
		os.MkdirAll(ppath, os.ModePerm)
	}
	d, err := yaml.Marshal(p)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(p.GetProjectFilename(), d, 0644)
	return err
}

// LoadProject based on name and LocalDirectory
func (p *Project) LoadProject() error {
	if content, err := ioutil.ReadFile(p.GetProjectFilename()); err == nil {
		pdata := Project{}
		yaml.Unmarshal(content, &pdata)
		p.PID = pdata.PID
		p.Environments = pdata.Environments
		p.Applications = pdata.Applications
	}
	return nil
}

// CurrentProject values
var CurrentProject *Project

func init() {
	apps := make(map[string]Application)
	envs := make(map[string]Environment)
	name := prefab.GenerateName(false)
	CurrentProject = &Project{Name: name, Environments: envs, Applications: apps}
}
