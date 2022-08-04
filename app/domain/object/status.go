package object

type (
	StatusID = int64

	// Account account
	Status struct {
		// The status ID
		S_id StatusID `json:"-" db:"id"`

		// The internal ID of the account
		U_id AccountID `json:"-" db:"account_id"`

		// content
		Content string `json:"content,omitempty" db:"content"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
