package controllers

import (
	"fmt"
	"github.com/CalmBit/capybara/serializers"
	"github.com/kataras/iris/mvc"
	"reflect"
	_ "time"
	"github.com/gobuffalo/pop"
	"github.com/CalmBit/capybara/models"
)

// temporary
var typeName = "Account"

var APIAccountSerializer = serializers.ConstructSerializer(func(v interface{}) error {
	if reflect.TypeOf(v).Name() == typeName {
		return nil
	} else {
		return fmt.Errorf("bad serialize on type %s expecting type %s", reflect.TypeOf(v).Name(), typeName)
	}
}).AddFieldAsString("ID").AddField("Username").AddMethod("Acct", "acct").AddField("DisplayName").AddField("Locked").AddField("CreatedAt").AddField("Note").AddField("URL").AddMethod("Avatar", "avatar").AddMethod("AvatarStatic", "avatar_static").AddMethod("Header", "header").AddMethod("HeaderStatic", "header_static").AddField("FollowersCount").AddField("FollowingCount").AddField("StatusesCount")


type AccountController struct{}

func (a *AccountController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *AccountController) GetBy(id int64) mvc.Result {
	tx, err := pop.Connect("development")
	if err != nil {
		return mvc.Response{
			ContentType: "application/json",
			Content:     []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
		}
	}
	var account models.Account
	err = tx.Find(&account, id)
	if err != nil {
		return mvc.Response{
			ContentType: "application/json",
			Content:     []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
		}
	}
	buf, err := APIAccountSerializer.SerializeToJSON(account)
	if err != nil {
		return mvc.Response{
			ContentType: "application/json",
			Content:     []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
		}
	}
	return mvc.Response{
		ContentType: "application/json",
		Content:     buf,
	}
}
