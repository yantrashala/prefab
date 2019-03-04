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
	Name string
	Type string
	Repo string
}

// GetBuildEnvironmentTypes returns a list of supported environment types
func GetBuildEnvironmentTypes() ([]Environment, error) {
	url := "https://raw.githubusercontent.com/yantrashala/prefab-config/master/environments.yaml"
	defaultTransport := http.DefaultTransport.(*http.Transport)

	// Create new Transport that ignores self-signed SSL TODO: find a better way
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
	m := make(map[string][]Environment)

	err = yaml.Unmarshal([]byte(e), &m)

	var environments []Environment
	for k := range m {
		if k == "build" {
			repos := m[k]
			for r := range repos {
				environments = append(environments, repos[r])
			}
		}
	}
	return environments, nil
}
