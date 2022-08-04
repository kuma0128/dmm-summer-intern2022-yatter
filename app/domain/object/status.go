package object

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
