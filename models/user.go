package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
	"net"
	"database/sql/driver"
	"fmt"
)

type UserIP net.IPAddr

type User struct {
	ID                     int   `json:"id" db:"id"`
	Email                  string    `json:"email" db:"email"`
	CreatedAt              time.Time `json:"created_at" db:"created_at"`
	UpdatedAt              time.Time `json:"updated_at" db:"updated_at"`
	EncryptedPassword      string    `json:"encrypted_password" db:"encrypted_password"`
	ResetPasswordToken     string    `json:"reset_password_token" db:"reset_password_token"`
	ResetPasswordSentAt    time.Time `json:"reset_password_sent_at" db:"reset_password_sent_at"`
	RememberCreatedAt      time.Time `json:"remember_created_at" db:"remember_created_at"`
	SignInCount            int   `json:"sign_in_count" db:"sign_in_count"`
	CurrentSignInAt        time.Time `json:"current_sign_in_at" db:"current_sign_in_at"`
	LastSignInAt           time.Time `json:"last_sign_in_at" db:"last_sign_in_at"`
	CurrentSignInIp        UserIP `json:"current_sign_in_ip" db:"current_sign_in_ip"`
	LastSignInIp           UserIP      `json:"last_sign_in_ip" db:"last_sign_in_ip"`
	Admin                  bool   `json:"admin" db:"admin"`
	ConfirmationToken      string    `json:"confirmation_token" db:"confirmation_token"`
	ConfirmedAt            time.Time `json:"confirmed_at" db:"confirmed_at"`
	ConfirmationSentAt     time.Time `json:"confirmation_sent_at" db:"confirmation_sent_at"`
	UnconfirmedEmail       string    `json:"unconfirmed_email" db:"unconfirmed_email"`
	Locale                 string    `json:"locale" db:"locale"`
	EncryptedOtpSecret     string    `json:"encrypted_otp_secret" db:"encrypted_otp_secret"`
	EncryptedOtpSecretIv   string    `json:"encrypted_otp_secret_iv" db:"encrypted_otp_secret_iv"`
	EncryptedOtpSecretSalt string    `json:"encrypted_otp_secret_salt" db:"encrypted_otp_secret_salt"`
	ConsumedTimestep       int   `json:"consumed_timestep" db:"consumed_timestep"`
	OtpRequiredForLogin    bool   `json:"otp_required_for_login" db:"otp_required_for_login"`
	LastEmailedAt          time.Time `json:"last_emailed_at" db:"last_emailed_at"`
	OtpBackupCodes         string    `json:"otp_backup_codes" db:"otp_backup_codes"`
	FilteredLanguages      string    `json:"filtered_languages" db:"filtered_languages"`
	AccountID              int64    `json:"account_id" db:"account_id" belongs_to:"account"`
	Disabled               bool   `json:"disabled" db:"disabled"`
	Moderator              bool   `json:"moderator" db:"moderator"`
	InviteID               int   `json:"invite_id" db:"invite_id"`
	RememberToken          string    `json:"remember_token" db:"remember_token"`
}

var SampleAccount = Account{
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
	MovedToAccountID:      AccountId{0},
	FeaturedCollectionURL: "test.notlocal/CalmBit/featured",
}

func NewUser() User {
	return User{
		ID:                     0,
		Email:                  "",
		CreatedAt:              time.Now(),
		UpdatedAt:              time.Now(),
		EncryptedPassword:      "",
		ResetPasswordToken:     "",
		ResetPasswordSentAt:    time.Time{},
		RememberCreatedAt:      time.Time{},
		SignInCount:            0,
		CurrentSignInAt:        time.Time{},
		LastSignInAt:           time.Time{},
		CurrentSignInIp:        UserIP{IP:net.ParseIP("127.0.0.1")},
		LastSignInIp:           UserIP{IP:net.ParseIP("127.0.0.1")},
		Admin:                  false,
		ConfirmationToken:      "",
		ConfirmedAt:            time.Time{},
		ConfirmationSentAt:     time.Time{},
		UnconfirmedEmail:       "",
		Locale:                 "",
		EncryptedOtpSecret:     "",
		EncryptedOtpSecretIv:   "",
		EncryptedOtpSecretSalt: "",
		ConsumedTimestep:       0,
		OtpRequiredForLogin:    false,
		LastEmailedAt:          time.Time{},
		OtpBackupCodes:         "",
		FilteredLanguages:      "",
		AccountID:              0,
		Disabled:               false,
		Moderator:              false,
		InviteID:               0,
		RememberToken:          "",
	}
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
		&validators.TimeIsPresent{Field: u.CreatedAt, Name: "CreatedAt"},
		&validators.TimeIsPresent{Field: u.UpdatedAt, Name: "UpdatedAt"},
		&validators.StringIsPresent{Field: u.EncryptedPassword, Name: "EncryptedPassword"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

func (u *User) CreateAccount(tx *pop.Connection, username string) error {
	var account = Account{
		ID:                    0,
		CreatedAt:             time.Time{},
		UpdatedAt:             time.Time{},
		Username:              "",
		Domain:                "",
		Secret:                "",
		PrivateKey:            "",
		PublicKey:             "",
		RemoteURL:             "",
		SalmonURL:             "",
		HubURL:                "",
		Note:                  "",
		DisplayName:           "",
		URI:                   "",
		URL:                   "",
		AvatarFileName:        "",
		AvatarContentType:     "",
		AvatarFileSize:        0,
		AvatarUpdatedAt:       time.Time{},
		HeaderFileName:        "",
		HeaderContentType:     "",
		HeaderFileSize:        0,
		HeaderUpdatedAt:       time.Time{},
		AvatarRemoteURL:       "",
		SubscriptionExpiresAt: time.Time{},
		Silenced:              false,
		Suspended:             false,
		Locked:                false,
		HeaderRemoteURL:       "",
		StatusesCount:         0,
		FollowersCount:        0,
		FollowingCount:        0,
		LastWebfingeredAt:     time.Time{},
		InboxURL:              "",
		OutboxURL:             "",
		SharedInboxURL:        "",
		FollowersURL:          "",
		Protocol:              0,
		Memorial:              false,
		MovedToAccountID:      AccountId{},
		FeaturedCollectionURL: "",
	}
	account.Username = username
	valid, err := tx.ValidateAndCreate(&account)
	if valid.HasAny() {
		return fmt.Errorf("%s", valid.Error())
	}
	if err != nil {
		return err
	}
	u.AccountID = account.ID
	return nil
}

func (ip UserIP) Value() (driver.Value, error) {
	return driver.Value(ip.IP.String()), nil
}

func (ip UserIP) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	}
	ip.IP = net.ParseIP(string(source))

	return nil
}