package model

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// Config holds the default configuration values for prefab
var Config map[interface{}]interface{}

// map[string]interface{}

// LoadConfig read the default configuration vales for prefab from the given filename
func LoadConfig(filename string) error {
	if content, err := ioutil.ReadFile(filename); err == nil {
		yaml.Unmarshal(content, &Config)
		//fmt.Printf("%+v", Config)
	} else {
		return err
	}
	return nil
}

// GetConfigAsYaml will return the serialized version of the glocal config
func GetConfigAsYaml() string {
	d, err := yaml.Marshal(&Config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return string(d)
}

// SaveConfig saved the updated configuration values to the specified file
func SaveConfig(filename string) error {
	return ioutil.WriteFile(filename, []byte(GetConfigAsYaml()), 0644)
}

func init() {
	Config = make(map[interface{}]interface{})
}
