package controllers

import (
	"github.com/kataras/iris/mvc"
	"github.com/CalmBit/capybara/serializers"
	"reflect"
	"fmt"
	"github.com/CalmBit/capybara/models"
	"time"
)

// temporary
var typeName = "Account"

var APIAccountSerializer = serializers.ConstructSerializer(func(v interface{}) error {
	if reflect.TypeOf(v).Name() == typeName {
		return nil
	} else {
		return fmt.Errorf("bad serialize on type %s expecting type %s", reflect.TypeOf(v).Name(), typeName)}
	}).AddFieldAsString("ID",
	).AddField("Username",
	).AddMethod("Acct", "acct",
	).AddField("DisplayName",
	).AddField("Locked",
	).AddField("CreatedAt",
	).AddField("Note",
	).AddField("URL",
	).AddMethod("Avatar", "avatar",
	).AddMethod("AvatarStatic", "avatar_static",
	).AddMethod("Header", "header",
	).AddMethod("HeaderStatic", "header_static",
	).AddField("FollowersCount",
	).AddField("FollowingCount",
	).AddField("StatusesCount")

var sampleAccount = models.Account{
	ID:                    8675309,
	CreatedAt:             time.Now(),
	UpdatedAt:             time.Now(),
	Username:              "CalmBit",
	Domain:                "test.notlocal",
	Secret:                "very secret",
	PrivateKey:            "private",
	PublicKey:             "public",
	RemoteURL:             "remote url thing",
	SalmonURL:             "salmon url thing",
	HubURL:                "hub url thing",
	Note:                  "\u00A7\"hello world\"",
	DisplayName:           "CalmBit",
	URI:                   "test.notlocal/CalmBit",
	URL:                   "test.notlocal/CalmBit",
	AvatarFileName:        "test.notlocal/CalmBit.png",
	AvatarContentType:     "image/png",
	AvatarFileSize:        1337,
	AvatarUpdatedAt:       time.Now(),
	HeaderFileName:        "test.notlocal/CalmBit_header.png",
	HeaderContentType:     "image/gif",
	HeaderFileSize:        7331,
	HeaderUpdatedAt:       time.Now(),
	AvatarRemoteURL:       "test.notlocal/CalmBit.png",
	SubscriptionExpiresAt: time.Now(),
	Silenced:              false,
	Suspended:             false,
	Locked:                true,
	HeaderRemoteURL:       "test.notlocal/CalmBit_header.png",
	StatusesCount:         100,
	FollowersCount:        200,
	FollowingCount:        300,
	LastWebfingeredAt:     time.Now(),
	InboxURL:              "test.notlocal/CalmBit/inbox",
	OutboxURL:             "test.notlocal/CalmBit/outbox",
	SharedInboxURL:        "test.notlocal/inbox",
	FollowersURL:          "test.notlocal/CalmBit/followers",
	Protocol:              0,
	Memorial:              false,
	MovedToAccountID:      0,
	FeaturedCollectionURL: "test.notlocal/CalmBit/featured",
}

type AccountController struct{}

func (a *AccountController) BeforeActivation(b mvc.BeforeActivation) {
}

func (a *AccountController) GetBy(id int64) mvc.Result {
	buf, err := APIAccountSerializer.SerializeToJSON(sampleAccount)
	if err != nil {
		return mvc.Response{
			ContentType: "application/json",
			Content: []byte(fmt.Sprintf("{\"error\": \"%s\"}", err.Error())),
		}
	}
	return mvc.Response{
		ContentType: "application/json",
		Content: buf,
	}
}
