package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/CalmBit/capybara/controllers"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const capybaraVersionString = "v0.1.0"

func main() {
	log.SetOutput(os.Stderr)
	textFormatter := new(prefixed.TextFormatter)
	textFormatter.FullTimestamp = true
	textFormatter.TimestampFormat = "Jan 01 2006 15:04:05"
	log.SetFormatter(textFormatter)
	log.SetLevel(log.DebugLevel)

	log.Infof("Capybara %s is starting up...", capybaraVersionString)
	log.Debugf("Establishing connection to database...")
	_, err := pop.Connect("development")
	if err != nil {
		log.Fatalf("Unable to establish database connection: %s", err.Error())
		return
	}
	log.Debugf("Starting up Iris....")
	app := iris.New()
	mvc.Configure(app.Party("/api/v1/accounts"), accounts)
	log.Infof("Listening on 8080...")

	app.Run(iris.Addr(":8080"), iris.WithoutStartupLog, iris.WithoutServerError(iris.ErrServerClosed))
	log.Infof("Goodbye! :)")
}

func accounts(app *mvc.Application) {
	app.Handle(new(controllers.AccountController))
}
