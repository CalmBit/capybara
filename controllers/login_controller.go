package controllers

import (
	"encoding/base64"
	"fmt"
	"github.com/CalmBit/capybara/middleware"
	"github.com/CalmBit/capybara/models"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/sec51/twofactor"
	"golang.org/x/crypto/bcrypt"
	"html"
)

type LoginController struct{}

func (l *LoginController) BeforeActivation(b mvc.BeforeActivation) {
}

func (l *LoginController) BeginRequest(ctx iris.Context) {
}

func (l *LoginController) EndRequest(ctx iris.Context) {
}

func (l *LoginController) Get(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	if s.Get("authenticated") != nil {
		return mvc.Response{
			Path: "/",
		}
	}
	return mvc.View{
		Name: "login.pug",
	}
}

func (l *LoginController) Post(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	email := html.EscapeString(ctx.FormValue("email"))

	tx, err := pop.Connect("development")
	if err != nil {
		s.SetFlash("error", err.Error())
		return mvc.Response{
			Path: "/login",
		}
	}

	var user models.User
	query := tx.Where(fmt.Sprintf("email = '%s'", email))
	err = query.First(&user)
	if err != nil {
		s.SetFlash("error", "Invalid username or password!")
		return mvc.Response{
			Path: "/login",
		}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.EncryptedPassword), []byte(ctx.FormValue("password")))
	if err != nil {
		s.SetFlash("error", "Invalid username or password!")
		return mvc.Response{
			Path: "/login",
		}
	} else {
		s.Set("user_id", user.ID)
		s.Set("pass_confirmed", true)
		if user.OtpRequiredForLogin {
			return mvc.Response{
				Path: "/otp/login",
			}
		}
		var acct models.Account
		err = tx.Find(&acct, user.AccountID)
		if err != nil {
			s.SetFlash("error", err.Error())
			s.Delete("user_id")
			s.Delete("pass_confirmed")
			return mvc.Response{
				Path: "/login",
			}
		}
		s.Set("account_id", acct.ID)
		s.Set("username", acct.Username)
		s.Set("email", user.Email)
		s.Set("display_name", acct.DisplayName)
		s.Set("authenticated", true)
		return mvc.Response{
			Path: "/",
		}
	}
}

func (l *LoginController) PostConfirm(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)

	if !s.GetBooleanDefault("pass_confirmed", false) {
		s.SetFlash("error", "An error occurred. Please try again.")
		return mvc.Response{
			Path: "/login",
		}
	}
	id, err := s.GetInt64("user_id")
	if err != nil {
		s.SetFlash("error", err.Error())
		s.Delete("pass_confirmed")
		return mvc.Response{
			Path: "/login",
		}
	}
	tx, err := pop.Connect("development")
	if err != nil {
		s.SetFlash("error", err.Error())
		s.Delete("pass_confirmed")
		s.Delete("user_id")
		return mvc.Response{
			Path: "/login",
		}
	}
	var user models.User
	tx.Find(&user, id)
	unmarshalSecret := make([]byte, 256)
	n, err := base64.StdEncoding.Decode(unmarshalSecret, []byte(user.EncryptedOtpSecret))
	unmarshalSecret = unmarshalSecret[:n]
	otp, err := twofactor.TOTPFromBytes(unmarshalSecret, "capybara")
	if err != nil {
		s.SetFlash("error", err.Error())
		s.Delete("pass_confirmed")
		s.Delete("user_id")
		return mvc.Response{
			Path: "/login",
		}
	}
	err = otp.Validate(ctx.FormValue("code_confirm"))
	if err != nil {
		s.SetFlash("error", err.Error())
		return mvc.Response{
			Path: "/otp/login",
		}
	}
	var acct models.Account
	err = tx.Find(&acct, user.AccountID)
	if err != nil {
		s.SetFlash("error", err.Error())
		s.Delete("pass_confirmed")
		s.Delete("user_id")
		return mvc.Response{
			Path: "/login",
		}
	}
	s.Set("user_id", user.ID)
	s.Set("account_id", acct.ID)
	s.Set("username", acct.Username)
	s.Set("email", user.Email)
	s.Set("display_name", acct.DisplayName)
	s.Set("authenticated", true)
	return mvc.Response{
		Path: "/",
	}
}
