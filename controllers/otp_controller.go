package controllers

import (
	"crypto"
	"encoding/base64"
	"fmt"
	"github.com/CalmBit/capybara/middleware"
	"github.com/CalmBit/capybara/models"
	"github.com/gobuffalo/pop"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/sec51/twofactor"
)

type OTPController struct{}

func (o *OTPController) BeforeActivation(b mvc.BeforeActivation) {
}

func (o *OTPController) BeginRequest(ctx iris.Context) {
}

func (o *OTPController) EndRequest(ctx iris.Context) {
}

func (o *OTPController) GetSettings(ctx iris.Context) {
	s := middleware.GetSession(ctx)
	if s.Get("authenticated") == nil {
		s.SetFlash("error", "You need to be logged in to do that.")
		ctx.Redirect("/login")
	} else {
		otp, err := twofactor.NewTOTP(s.GetString("email"), "capybara", crypto.SHA1, 6)
		if err != nil {
			s.SetFlash("error", "There was a problem processing your request.")
			ctx.Redirect("/about")
		} else {
			qrpng, err := otp.QR()
			if err != nil {
				s.SetFlash("error", "There was a problem processing the QR code.")
				ctx.Redirect("/about")
			} else {
				var qrbuffer = make([]byte, ((len(qrpng)+2)/3)*4)
				base64.StdEncoding.Encode(qrbuffer, qrpng)
				s.SetFlash("otp", otp)
				ctx.ViewData("qr_code", fmt.Sprintf("data:image/png;base64, %s", string(qrbuffer)))
				ctx.View("otp.pug")
			}
		}

	}
}

func (o *OTPController) GetLogin(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	if s.Get("authenticated") != nil {
		return mvc.Response{
			Path: "/",
		}
	}
	if !s.GetBooleanDefault("pass_confirmed", false) {
		s.SetFlash("error", "An error occurred. Please try again.")
		return mvc.Response{
			Path: "/login",
		}
	}
	return mvc.View{
		Name: "otp_login.pug",
	}

}

func (o *OTPController) Post(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	otp := s.GetFlash("otp").(*twofactor.Totp)
	err := otp.Validate(ctx.FormValue("code_confirm"))
	if err != nil {
		s.SetFlash("error", "There was a problem processing the confirmation code")
		return mvc.Response{
			Path: "/otp/settings",
		}
	}
	tx, err := pop.Connect("development")
	if err != nil {
		s.SetFlash("error", "The code was confirmed, but we had trouble storing it. Please try again.")
		return mvc.Response{
			Path: "/otp/settings",
		}
	}
	var user models.User
	id, err := s.GetInt64("user_id")
	if err != nil {
		s.SetFlash("error", "You need to be logged in to do that.")
		return mvc.Response{
			Path: "/login",
		}
	}
	tx.Find(&user, id)
	buf, err := otp.ToBytes()
	if err != nil {
		s.SetFlash("error", "Unable to serialize secret? (this is bad)")
		return mvc.Response{
			Path: "/otp/settings",
		}
	}
	bufencode := make([]byte, ((len(buf)+2)/3)*4)
	base64.StdEncoding.Encode(bufencode, buf)
	user.EncryptedOtpSecret = string(bufencode)
	user.OtpRequiredForLogin = true
	valid, err := tx.ValidateAndSave(&user)
	if valid.HasAny() {
		s.SetFlash("error", valid.Error())
		return mvc.Response{
			Path: "/otp/settings",
		}
	} else if err != nil {
		s.SetFlash("error", err.Error())
		return mvc.Response{
			Path: "/otp/settings",
		}
	}
	s.SetFlash("error", "Success!")
	return mvc.Response{
		Path: "/about",
	}
}
