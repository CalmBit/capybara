package middleware

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var session = sessions.New(sessions.Config{Cookie: "capybara_session"})

func SetupSessionBackend(db sessions.Database) {
	session.UseDatabase(db)
}
func InitializeSession(ctx iris.Context) {
	NewSession(ctx)
	ctx.Next()
}

func NewSession(ctx iris.Context) *sessions.Session {
	s := session.Start(ctx)
	ctx.Values().Set("currentSession", s)
	ctx.ViewData("error", s.GetFlashString("error"))
	ctx.ViewData("success", s.GetFlashString("success"))
	return s
}

func GetSession(ctx iris.Context) *sessions.Session {
	return ctx.Values().Get("currentSession").(*sessions.Session)
}

func KillSession(ctx iris.Context) {
	s := GetSession(ctx)
	s.Delete("authenticated")
	s.Delete("user_id")
	s.Delete("pass_confirmed")
	s.Delete("account_id")
	s.Delete("username")
	s.Delete("email")
	s.Delete("display_name")
	s.Delete("authenticated")
}
