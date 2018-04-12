package controllers

import (
	"fmt"
	"github.com/CalmBit/capybara/models"
	"github.com/kataras/iris/mvc"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/kataras/iris"
	"github.com/CalmBit/capybara/middleware"
)

type PublicController struct{}

func (a *PublicController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *PublicController) BeginRequest(ctx iris.Context) {
}

func (a *PublicController) EndRequest(ctx iris.Context) {
}

func (a *PublicController) Get(ctx iris.Context) mvc.Result {
	s := middleware.GetSession(ctx)
	if s.Get("authenticated") == nil {
		return mvc.Response{
			Path: "/about",
		}
	}
	return mvc.View{
		Name: "index.pug",
	}
}

func (a *PublicController) GetBy(ctx iris.Context, param string) mvc.Result {
	tx, err := pop.Connect("development")
	if err != nil {
		return mvc.Response{
			Code: 500,
		}
	}
	id, err := uuid.FromString(param)
	if err != nil {
		return mvc.Response{
			Code: 404,
		}
	}
	var account models.Account
	err = tx.Where(fmt.Sprintf("uuid = '%s'", id.String())).First(&account)
	if err != nil {
		return mvc.Response{
			Code: 404,
		}
	}

	if !account.Local() {
		return mvc.Response{
			Code: 404,
		}
	}

	ctx.ViewData("account", account)

	var statuses []models.Status
	err = tx.Where(fmt.Sprintf("account_id = %d ORDER BY created_at DESC", account.ID)).All(&statuses)
	if err != nil {
		return mvc.Response{
			Code: 500,
			Content: []byte(err.Error()),
		}
	}

	for i, status := range statuses {
		err = tx.Where(fmt.Sprintf("id = '%d'", status.AccountID)).First(&statuses[i].StatusAccount)
		if err != nil {
			fmt.Printf("%s", err.Error())
		} else {
			fmt.Printf("%s\n", status.StatusAccount.Username)
		}
	}

	fmt.Printf("%s\n", statuses[0].StatusAccount.Username)

	ctx.ViewData("statuses", statuses)




	return mvc.View{
		Name: "public_user.pug",
	}
}

