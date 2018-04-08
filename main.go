package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/x-cray/logrus-prefixed-formatter"

	"github.com/CalmBit/capybara/controllers"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
)

const capybaraVersionString = "v0.1.0"

func main() {

	log.SetOutput(os.Stderr)
	textFormatter := new(prefixed.TextFormatter)
	textFormatter.FullTimestamp = true
	textFormatter.TimestampFormat = "Jan 01 2006 15:04:05"
	log.SetFormatter(textFormatter)
	log.SetLevel(log.DebugLevel)
	controllers.LoadSettings()
	log.Infof("Capybara %s is starting up...", capybaraVersionString)
	log.Debugf("Establishing connection to database...")
	tx, err := pop.Connect("development")
	if err != nil {
		log.Fatalf("Unable to establish database connection: %s", err.Error())
		return
	}
	log.Debugf("Got database connection %s", tx.ID)
	log.Debugf("Establishing connection to redis...")
	cache := redis.New(service.Config{
		Network:     service.DefaultRedisNetwork,
		Addr:        service.DefaultRedisAddr,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: service.DefaultRedisIdleTimeout,
		Prefix:      "",
	})
	iris.RegisterOnInterrupt(func() {
		cache.Close()
		tx.Close()
	})

	log.Debugf("Starting up Iris....")
	app := iris.New()
	app.StaticWeb("/public/", "./public")
	pugEngine := iris.Pug("./views", ".pug")
	pugEngine.Reload(true)
	app.RegisterView(pugEngine)
	controllers.Session.UseDatabase(cache)
	app.Get("/", func(ctx iris.Context) {
		s := controllers.Session.Start(ctx)
		if s.Get("user_id") == nil {
			ctx.Redirect("/about")
		}
	})
	mvc.Configure(app.Party("/api/v1/accounts"), accounts)
	mvc.Configure(app.Party("/register"), registrations)
	mvc.Configure(app.Party("/login"), logins)
	mvc.Configure(app.Party("/about"), about)
	mvc.Configure(app.Party("/otp"), otp)
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

func logins(app *mvc.Application) {
	app.Handle(new(controllers.LoginController))
}

func about(app *mvc.Application) {
	app.Handle(new(controllers.AboutController))
}

func otp(app *mvc.Application) {
	app.Handle(new(controllers.OTPController))
}