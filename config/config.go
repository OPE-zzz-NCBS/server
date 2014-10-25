package config

import (
	"os"
	"encoding/json"
)

type Server struct {
	Port int
}

type Database struct {
	Host string
	Username string
	Password string
	Name string
}

type Configuration struct {
	Server *Server
	Database *Database
}

var configuration *Configuration

func Get() (*Configuration, error) {
	if configuration != nil {
		return configuration, nil
	}

	file, err := os.Open("conf.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration = new(Configuration)
	err = decoder.Decode(configuration)
	if err != nil {
		configuration = nil
		return nil, err
	}
	return configuration, nil
}

