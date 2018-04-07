package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/CalmBit/capybara/controllers"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const capybaraVersionString = "v0.1.0"

// Temporary, offload somewhere else, config file?
type Settings struct {
	RegistrationsOpen bool `yaml:"registrationsOpen"`
}

func main() {
	var settings Settings
	setbuff, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("unable to read config file - %s", err.Error())
		return
	}
	err = yaml.Unmarshal(setbuff, &settings)
	if err != nil {
		log.Fatalf("unable to unmarshal config file - %s", err.Error())
		return
	}
	log.SetOutput(os.Stderr)
	textFormatter := new(prefixed.TextFormatter)
	textFormatter.FullTimestamp = true
	textFormatter.TimestampFormat = "Jan 01 2006 15:04:05"
	log.SetFormatter(textFormatter)
	log.SetLevel(log.DebugLevel)

	log.Infof("Capybara %s is starting up...", capybaraVersionString)
	log.Debugf("Establishing connection to database...")
	tx, err := pop.Connect("development")
	defer tx.Close()
	if err != nil {
		log.Fatalf("Unable to establish database connection: %s", err.Error())
		return
	}
	log.Debugf("Got database connection %s", tx.ID)
	log.Debugf("Starting up Iris....")
	app := iris.New()
	app.StaticWeb("/public/", "./public")
	pugEngine := iris.Pug("./views", ".pug")
	pugEngine.Reload(true)
	app.RegisterView(pugEngine)

	app.Get("/", func(ctx iris.Context) {
		ctx.ViewData("settings", settings)
		ctx.View("index.pug")
	})
	mvc.Configure(app.Party("/api/v1/accounts"), accounts)
	mvc.Configure(app.Party("/register"), registrations)
	log.Infof("Listening on 8080...")


	app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	log.Infof("Goodbye! :)")
}

func accounts(app *mvc.Application) {
	app.Handle(new(controllers.AccountController))
}

func registrations(app *mvc.Application) {
	app.Handle(new(controllers.RegistrationController))
}