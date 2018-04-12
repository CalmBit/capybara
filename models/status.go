package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
	"fmt"
	"html/template"
	"github.com/CalmBit/capybara/util"
)

type Status struct {
	ID                 int64     `json:"id" db:"id"`
	URI                string    `json:"uri" db:"uri"`
	Text               string    `json:"text" db:"text"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
	InReplyToID        int       `json:"in_reply_to_id" db:"in_reply_to_id"`
	ReblogOfID         int       `json:"reblog_of_id" db:"reblog_of_id"`
	URL                string    `json:"url" db:"url"`
	Sensitive          bool      `json:"sensitive" db:"sensitive"`
	Visibility         int       `json:"visibility" db:"visibility"`
	SpoilerText        string    `json:"spoiler_text" db:"spoiler_text"`
	Reply              bool      `json:"reply" db:"reply"`
	FavouritesCount    int       `json:"favourites_count" db:"favourites_count"`
	ReblogsCount       int       `json:"reblogs_count" db:"reblogs_count"`
	Language           string    `json:"language" db:"language"`
	ConversationID     int       `json:"conversation_id" db:"conversation_id"`
	Local              bool      `json:"local" db:"local"`
	AccountID          int       `json:"account_id" db:"account_id"`
	ApplicationID      int       `json:"application_id" db:"application_id"`
	InReplyToAccountID int       `json:"in_reply_to_account_id" db:"in_reply_to_account_id"`
	StatusAccount	Account		 `db:"-"`

}

// String is not required by pop and may be deleted
func (s Status) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Statuses is not required by pop and may be deleted
type Statuses []Status

// String is not required by pop and may be deleted
func (s Statuses) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Status) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.URI, Name: "URI"},
		&validators.StringIsPresent{Field: s.Text, Name: "Text"},
		&validators.TimeIsPresent{Field: s.CreatedAt, Name: "CreatedAt"},
		&validators.TimeIsPresent{Field: s.UpdatedAt, Name: "UpdatedAt"},
		&validators.StringIsPresent{Field: s.URL, Name: "URL"},
		&validators.StringIsPresent{Field: s.SpoilerText, Name: "SpoilerText"},
		&validators.StringIsPresent{Field: s.Language, Name: "Language"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Status) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Status) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (s *Status) TimeAgo() string {
	return util.Ago(s.CreatedAt)
}

func (s *Status) Avatar() template.HTML {
	tx, err := pop.Connect("development")
	if err != nil {
		return "(error - couldn't connect to database)"
	}
	var account Account
	err = tx.Where(fmt.Sprintf("id = '%d'", s.AccountID)).First(&account)
	if err != nil {
		return "(error - couldn't find user)"
	}

	if account.AvatarFileName == "" {
		return template.HTML(fmt.Sprintf("<img id=\"status_img\" src=\"%s\" height=\"48\" width=\"48\"/>", "public/img/avatar/missing.png"))

	}
	return template.HTML(fmt.Sprintf("<img id=\"status_img\" src=\"%s\"/>", account.AvatarFileName))
}