package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/CalmBit/capybara/models"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/gobuffalo/pop"
)

type RegistrationController struct{}

func (a *RegistrationController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *RegistrationController) Post(ctx iris.Context) mvc.Result {
	s := Session.Start(ctx)
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
		err = newUser.CreateAccount(tx, ctx.FormValue("username"))
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
		buf, err := json.Marshal(newUser)
		if err != nil {
			s.SetFlash("error", err.Error())
			return mvc.Response{
				Path: "/about",
			}
		}
		return mvc.Response{
			ContentType: "application/json",
			Content:     buf,
		}
	}

	s.SetFlash("error", "Passwords did not match!")
	return mvc.Response{
		Path: "/about",
	}
}
