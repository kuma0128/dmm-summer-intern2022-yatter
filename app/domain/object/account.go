package object

import (
	"fmt"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	AccountID    = int64
	PasswordHash = string

	// Account account
	Account struct {
		// The internal ID of the account
		ID AccountID `json:"id" db:"id"`

		// The username of the account
		Username string `json:"username,omitempty" db:"username"`

		// The username of the account
		PasswordHash string `json:"-" db:"password_hash"`

		// The account's display name
		DisplayName *string `json:"display_name,omitempty" db:"display_name"`

		// URL to the avatar image
		Avatar *string `json:"avatar,omitempty"`

		// URL to the header image
		Header *string `json:"header,omitempty"`

		// Biography of user
		Note *string `json:"note,omitempty"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`

		//The number of followers for the account
		Followers_count int32 `json:"followers.count"`

		//The number of accounts the given account is following
		Following_count int32 `json:"following_count"`
	}

	//relationship
	Relationship struct {
		ID AccountID `json:"id"`

		//Whether the user is currently following the account
		Following bool `json:"following" db:"follower_id"`

		//Whether the user is currently being followed by the account
		Followed_by bool `json:"followed_by" db:"followee_id"`
	}
)

// Create new account object
func CreateAccountobject(username, password string) (*Account, error) {
	if len(username) > 10 {
		return &Account{}, errors.New("username is too long")
	}
	account := Account{
		Username: username,
	}
	if err := account.SetPassword(password); err != nil {
		return &Account{}, errors.New("password is too short")
	}
	return &account, nil

}

// Check if given password is match to account's password
func (a *Account) CheckPassword(pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(a.PasswordHash), []byte(pass)) == nil
}

// Hash password and set it to account object
func (a *Account) SetPassword(pass string) error {
	if len(pass) < 5 {
		return errors.New("password is too short")
	}
	passwordHash, err := generatePasswordHash(pass)
	if err != nil {
		return fmt.Errorf("generate error: %w", err)
	}
	a.PasswordHash = passwordHash
	return nil
}

func generatePasswordHash(pass string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hashing password failed: %w", errors.WithStack(err))
	}
	return PasswordHash(hash), nil
}
