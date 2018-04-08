package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AboutController struct{}

func (a *AboutController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *AboutController) BeginRequest(ctx iris.Context) {
	s := Session.Start(ctx)
	ctx.ViewData("error", s.GetFlashString("error"))
	ctx.ViewData("settings", GlobalSettings)
}

func (a *AboutController) EndRequest(ctx iris.Context) {
}

func (a *AboutController) Get(ctx iris.Context) {
	s := Session.Start(ctx)
	if s.Get("username") != nil {
		ctx.ViewData("username", s.Get("username"))
	}
	ctx.View("about.pug")
}

