package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"fmt"
)

type Account struct {
	ID                    int64     `json:"id" db:"id"`
	CreatedAt             time.Time `json:"created_at" db:"created_at"`
	UpdatedAt             time.Time `json:"updated_at" db:"updated_at"`
	Username              string    `json:"username" db:"username"`
	Domain                string    `json:"domain" db:"domain"`
	Secret                string    `json:"secret" db:"secret"`
	PrivateKey            string    `json:"private_key" db:"private_key"`
	PublicKey             string    `json:"public_key" db:"public_key"`
	RemoteURL             string    `json:"remote_url" db:"remote_url"`
	SalmonURL             string    `json:"salmon_url" db:"salmon_url"`
	HubURL                string    `json:"hub_url" db:"hub_url"`
	Note                  string    `json:"note" db:"note"`
	DisplayName           string    `json:"display_name" db:"display_name"`
	URI                   string    `json:"uri" db:"uri"`
	URL                   string    `json:"url" db:"url"`
	AvatarFileName        string    `json:"avatar_file_name" db:"avatar_file_name"`
	AvatarContentType     string    `json:"avatar_content_type" db:"avatar_content_type"`
	AvatarFileSize        int32     `json:"avatar_file_size" db:"avatar_file_size"`
	AvatarUpdatedAt       time.Time `json:"avatar_updated_at" db:"avatar_updated_at"`
	HeaderFileName        string    `json:"header_file_name" db:"header_file_name"`
	HeaderContentType     string    `json:"header_content_type" db:"header_content_type"`
	HeaderFileSize        int32     `json:"header_file_size" db:"header_file_size"`
	HeaderUpdatedAt       time.Time `json:"header_updated_at" db:"header_updated_at"`
	AvatarRemoteURL       string    `json:"avatar_remote_url" db:"avatar_remote_url"`
	SubscriptionExpiresAt time.Time `json:"subscription_expires_at" db:"subscription_expires_at"`
	Silenced              bool      `json:"silenced" db:"silenced"`
	Suspended             bool      `json:"suspended" db:"suspended"`
	Locked                bool      `json:"locked" db:"locked"`
	HeaderRemoteURL       string    `json:"header_remote_url" db:"header_remote_url"`
	StatusesCount         int32     `json:"statuses_count" db:"statuses_count"`
	FollowersCount        int32     `json:"followers_count" db:"followers_count"`
	FollowingCount        int32     `json:"following_count" db:"following_count"`
	LastWebfingeredAt     time.Time `json:"last_webfingered_at" db:"last_webfingered_at"`
	InboxURL              string    `json:"inbox_url" db:"inbox_url"`
	OutboxURL             string    `json:"outbox_url" db:"outbox_url"`
	SharedInboxURL        string    `json:"shared_inbox_url" db:"shared_inbox_url"`
	FollowersURL          string    `json:"followers_url" db:"followers_url"`
	Protocol              int32     `json:"protocol" db:"protocol"`
	Memorial              bool      `json:"memorial" db:"memorial"`
	MovedToAccountID      int64     `json:"moved_to_account_id" db:"moved_to_account_id"`
	FeaturedCollectionURL string    `json:"featured_collection_url" db:"featured_collection_url"`
}

// String is not required by pop and may be deleted
func (a Account) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Accounts is not required by pop and may be deleted
type Accounts []Account

// String is not required by pop and may be deleted
func (a Accounts) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Account) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: a.Username, Name: "Username"},
		&validators.StringIsPresent{Field: a.Domain, Name: "Domain"},
		&validators.StringIsPresent{Field: a.Secret, Name: "Secret"},
		&validators.StringIsPresent{Field: a.PrivateKey, Name: "PrivateKey"},
		&validators.StringIsPresent{Field: a.PublicKey, Name: "PublicKey"},
		&validators.StringIsPresent{Field: a.RemoteURL, Name: "RemoteURL"},
		&validators.StringIsPresent{Field: a.SalmonURL, Name: "SalmonURL"},
		&validators.StringIsPresent{Field: a.HubURL, Name: "HubURL"},
		&validators.TimeIsPresent{Field: a.CreatedAt, Name: "CreatedAt"},
		&validators.TimeIsPresent{Field: a.UpdatedAt, Name: "UpdatedAt"},
		&validators.StringIsPresent{Field: a.Note, Name: "Note"},
		&validators.StringIsPresent{Field: a.DisplayName, Name: "DisplayName"},
		&validators.StringIsPresent{Field: a.URI, Name: "URI"},
		&validators.StringIsPresent{Field: a.URL, Name: "URL"},
		&validators.StringIsPresent{Field: a.AvatarFileName, Name: "AvatarFileName"},
		&validators.StringIsPresent{Field: a.AvatarContentType, Name: "AvatarContentType"},
		&validators.TimeIsPresent{Field: a.AvatarUpdatedAt, Name: "AvatarUpdatedAt"},
		&validators.StringIsPresent{Field: a.HeaderFileName, Name: "HeaderFileName"},
		&validators.StringIsPresent{Field: a.HeaderContentType, Name: "HeaderContentType"},
		&validators.TimeIsPresent{Field: a.HeaderUpdatedAt, Name: "HeaderUpdatedAt"},
		&validators.StringIsPresent{Field: a.AvatarRemoteURL, Name: "AvatarRemoteURL"},
		&validators.TimeIsPresent{Field: a.SubscriptionExpiresAt, Name: "SubscriptionExpiresAt"},
		&validators.StringIsPresent{Field: a.HeaderRemoteURL, Name: "HeaderRemoteURL"},
		&validators.TimeIsPresent{Field: a.LastWebfingeredAt, Name: "LastWebfingeredAt"},
		&validators.StringIsPresent{Field: a.InboxURL, Name: "InboxURL"},
		&validators.StringIsPresent{Field: a.OutboxURL, Name: "OutboxURL"},
		&validators.StringIsPresent{Field: a.SharedInboxURL, Name: "SharedInboxURL"},
		&validators.StringIsPresent{Field: a.FollowersURL, Name: "FollowersURL"},
		&validators.StringIsPresent{Field: a.FeaturedCollectionURL, Name: "FeaturedCollectionURL"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Account) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Account) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (a Account) Local() bool {
	return a.Domain == ""
}

func (a Account) Acct() string {
	if a.Local() {
		return a.Username
	} else {
		return fmt.Sprintf("%s@%s", a.Username, a.Domain)
	}
}

func (a Account) Avatar() string {
	return "avatar_placeholder"
}

func (a Account ) AvatarStatic() string {
	return "static_avatar_placeholder"
}

func (a Account) Header() string {
	return "header_placeholder"
}

func (a Account ) HeaderStatic() string {
	return "header_avatar_placeholder"
}