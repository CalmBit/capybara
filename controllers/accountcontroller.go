package controllers

import (
	"github.com/kataras/iris/mvc"
)

type AccountController struct{}

func (a *AccountController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *AccountController) GetBy(id int64) mvc.Result {
	return mvc.Response{
		ContentType: "application/json",
	}
}
