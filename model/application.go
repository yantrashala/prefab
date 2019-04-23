package model

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Application describes a template of the app repo
type Application struct {
	Name           string
	Type           string
	Repo           string
	Summary        string
	Config         map[string]string
	LocalDirectory string
	SCM
}

func getApplicationTypes() (map[string][]Application, error) {

	url := "https://raw.githubusercontent.com/yantrashala/prefab-config/master/apps.yaml"
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
	m := make(map[string][]Application)

	err = yaml.Unmarshal([]byte(e), &m)

	return m, err
}

// GetApplicationTypes returns a list of applications types
func GetApplicationTypes() ([]string, error) {
	m, err := getApplicationTypes()
	var appTypes []string
	for k := range m {
		appTypes = append(appTypes, k)
	}
	return appTypes, err
}

// GetApplications returns a list of applications of the specified types
func GetApplications(appType string) []Application {
	var apps []Application
	m, err := getApplicationTypes()
	if err != nil {
		log.Fatal(err)
	}

	for k := range m {
		if k == appType {
			repos := m[k]
			for r := range repos {
				repos[r].Config = make(map[string]string)
				apps = append(apps, repos[r])
			}
		}
	}
	return apps
}
