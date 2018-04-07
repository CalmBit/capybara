package controllers

import (
	"fmt"
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
	if ctx.FormValue("password") == ctx.FormValue("password_confirm") {
		passhash, err := bcrypt.GenerateFromPassword([]byte(ctx.FormValue("password")), 14)
		if err != nil {
			return mvc.Response{
				ContentType: "application/json",
				Content:  []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
			}
		}
		new_user := models.NewUser()
		new_user.Email = ctx.FormValue("email")
		new_user.EncryptedPassword = string(passhash)
		tx, err := pop.Connect("development")
		if err != nil {
			return mvc.Response{
				ContentType: "application/json",
				Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
			}
		}
		err = new_user.CreateAccount(tx, ctx.FormValue("username"))
		if err != nil {
			return mvc.Response{
				ContentType: "application/json",
				Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
			}
		}
		validate, err := tx.ValidateAndCreate(&new_user)
		if validate.HasAny() {
			return mvc.Response{
				ContentType: "application/json",
				Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", validate.Error())),
			}
		} else if err != nil {
			return mvc.Response{
				ContentType: "application/json",
				Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
			}
		}
		buf, err := json.Marshal(new_user)
		if err != nil {
			return mvc.Response{
				ContentType: "application/json",
				Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
			}
		}
		return mvc.Response{
			ContentType: "application/json",
			Content:     buf,
		}
	}

	return mvc.Response{
		ContentType: "application/json",
		Content:     []byte(fmt.Sprintf("{\"error\": \"passwords did not match\"}")),
	}
}
