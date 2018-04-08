package controllers

import (
	"github.com/kataras/iris/sessions"
	"github.com/kataras/iris"
)

var Session = sessions.New(sessions.Config{Cookie: "capybara_session"})

func Invalidate(ctx iris.Context) {
	Session.Destroy(ctx)
}