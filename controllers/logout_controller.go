package controllers

import (
	"github.com/CalmBit/capybara/middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type LogoutController struct{}

func (l *LogoutController) BeforeActivation(b mvc.BeforeActivation) {
}

func (l *LogoutController) BeginRequest(ctx iris.Context) {
}

func (l *LogoutController) EndRequest(ctx iris.Context) {
}

func (l *LogoutController) Post(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)

	if s.Get("authenticated") == nil {
		s.SetFlash("error", "You need to be logged in to log out!")
		return mvc.Response{
			Path: "/login",
		}
	}

	middleware.KillSession(ctx)

	s = middleware.NewSession(ctx)

	s.SetFlash("success", "You have been logged out.")

	return mvc.Response{
		Path: "/about",
	}
}
