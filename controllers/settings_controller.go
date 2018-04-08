package controllers

import (
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

type Settings struct {
	RegistrationsOpen bool `yaml:"registrationsOpen"`
}

var GlobalSettings Settings

func LoadSettings() {
	setbuff, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("unable to read config file - %s", err.Error())
		return
	}
	err = yaml.Unmarshal(setbuff, &GlobalSettings)
	if err != nil {
		log.Fatalf("unable to unmarshal config file - %s", err.Error())
		return
	}
}