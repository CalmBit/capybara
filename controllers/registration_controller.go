package controllers

import (
	"github.com/CalmBit/capybara/middleware"
	"github.com/CalmBit/capybara/models"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type RegistrationController struct{}

func (a *RegistrationController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *RegistrationController) BeginRequest(ctx iris.Context) {
}

func (a *RegistrationController) EndRequest(ctx iris.Context) {
}

func (a *RegistrationController) Post(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	if ctx.FormValue("password") == ctx.FormValue("password_confirm") {
		passhash, err := bcrypt.GenerateFromPassword([]byte(ctx.FormValue("password")), 14)
		if err != nil {
			s.SetFlash("error", err.Error())
			return mvc.Response{
				Path: "/about",
			}
		}
		newUser := models.NewUser()
		newUser.Email = ctx.FormValue("email")
		newUser.EncryptedPassword = string(passhash)
		tx, err := pop.Connect("development")
		if err != nil {
			s.SetFlash("error", err.Error())
			return mvc.Response{
				Path: "/about",
			}
		}
		acct, err := newUser.CreateAccount(tx, ctx.FormValue("username"))
		if err != nil {
			s.SetFlash("error", err.Error())
			return mvc.Response{
				Path: "/about",
			}
		}
		validate, err := tx.ValidateAndCreate(&newUser)
		if validate.HasAny() {
			s.SetFlash("error", validate.Error())
			return mvc.Response{
				Path: "/about",
			}
		} else if err != nil {
			s.SetFlash("error", err.Error())
			return mvc.Response{
				Path: "/about",
			}
		}

		return mvc.Response{
			Path: fmt.Sprintf("/%s", acct.UUID.String()),
		}
	}

	s.SetFlash("error", "Passwords did not match!")
	return mvc.Response{
		Path: "/about",
	}
}
