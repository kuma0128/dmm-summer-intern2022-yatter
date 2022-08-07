package object

import "errors"

type (
	StatusID = int64

	// Account account
	Status struct {
		// The status ID
		Sid StatusID `json:"id" db:"id"`

		// The accout id
		AccountID AccountID `json:"account_id" db:"account_id"`

		// The information of the account
		Account Account `json:"account" db:"account"`

		// content
		Content string `json:"content,omitempty" db:"content"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)

func CreateStatusobject(content string, account *Account) (*Status, error) {
	if len(content) > 120 {
		return &Status{}, errors.New("status content is too long")
	}
	status := Status{
		AccountID: account.ID,
		Account:   *account,
		Content:   content,
	}
	return &status, nil
}
