package config

import (
	"os"
	"encoding/json"
)

type Database struct {
	Host string
	Username string
	Password string
	Name string
}

type Configuration struct {
	Database Database
}

func (c *Configuration) Read() (err error) {
	file, err := os.Open("conf.json")
	if err != nil {
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		return
	}
	err = nil
	return
}

