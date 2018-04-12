package controllers

import (
	"github.com/CalmBit/capybara/middleware"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AboutController struct{}

func (a *AboutController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *AboutController) BeginRequest(ctx iris.Context) {
}

func (a *AboutController) EndRequest(ctx iris.Context) {
}

func (a *AboutController) Get(ctx iris.Context) {
	s := middleware.GetSession(ctx)
	if s.Get("username") != nil {
		ctx.ViewData("username", s.Get("username"))
	}
	ctx.View("about.pug")
}
