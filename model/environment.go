package model

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Environment describes a build or a run environment
type Environment struct {
	Name           string
	Type           string
	Repo           string
	Summary        string
	LocalDirectory string
	Config         map[string]string
	SCM
}

func getEnvironmentTypes() (map[string][]Environment, error) {

	url := "https://raw.githubusercontent.com/yantrashala/prefab-config/master/environments.yaml"
	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL
	// TODO: find a better way
	tr := &http.Transport{
		Proxy:                 defaultTransport.Proxy,
		DialContext:           defaultTransport.DialContext,
		MaxIdleConns:          defaultTransport.MaxIdleConns,
		IdleConnTimeout:       defaultTransport.IdleConnTimeout,
		ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
		TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	res, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	e, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// TODO: read into generic struct (m := make(map[interface{}]interface{})) and then transform to env
	m := make(map[string][]Environment)

	err = yaml.Unmarshal([]byte(e), &m)

	return m, err
}

// GetBuildEnvironmentTypes returns a list of supported environment types
func GetBuildEnvironmentTypes() ([]Environment, error) {
	m, err := getEnvironmentTypes()
	var environments []Environment
	for k := range m {
		if k == "build" {
			repos := m[k]
			for r := range repos {
				repos[r].Config = make(map[string]string)
				environments = append(environments, repos[r])
			}
		}
	}
	return environments, err
}

// GetRuntimeEnvironmentTypes returns a list of supported environment types
func GetRuntimeEnvironmentTypes() ([]Environment, error) {
	m, err := getEnvironmentTypes()
	var environments []Environment
	for k := range m {
		if k == "runtime" {
			repos := m[k]
			for r := range repos {
				repos[r].Config = make(map[string]string)
				environments = append(environments, repos[r])
			}
		}
	}
	return environments, err
}
