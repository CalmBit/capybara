package middleware

import (
	"github.com/kataras/iris"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Settings struct {
	RegistrationsOpen bool `yaml:"registrationsOpen"`
	URL string `yaml:"url"`
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

func InitializeSettings(ctx iris.Context) {
	ctx.ViewData("settings", GlobalSettings)
	ctx.Next()
}