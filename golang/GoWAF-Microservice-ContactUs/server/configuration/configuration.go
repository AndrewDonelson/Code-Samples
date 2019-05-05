package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//ServiceConfig holds data necessary for configuring application
type ServiceConfig struct {
	Database Database `json:"database"`
	General  General  `json:"general"`
}

//Database holds data for database configuring
type Database struct {
	Name     string `json:"db_name"`
	Password string `json:"password"`
	DbType   string `json:"db_type"`
}

type General struct {
	ServiceName string `json:"service_name"`
	Version     string `json:"version"`
}

//Extract - returns ServiceConfig struct
func Extract(path string) (*ServiceConfig, error) {
	conf := new(ServiceConfig)

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file, %s", err)
	}

	if err := json.Unmarshal(bytes, &conf); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	return conf, nil
}
