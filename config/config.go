package config

import (
	"os"
	"encoding/json"
)

type Database struct {
	Host string
	Username string
	Password string
}

type Configuration struct {
	Database Database
}

func (c *Configuration) Read() {
	file, err := os.Open("conf.json")
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(c)
	if err != nil {
		panic(err)
	}
}

