package object

type (
	StatusID = int64

	// Account account
	Status struct {
		// The status ID
		S_id StatusID `json:"-"`

		// The internal ID of the account
		U_id AccountID `json:"-"`

		// The username of the account
		Username string `json:"username,omitempty"`

		// URL to the header image
		Header *string `json:"header,omitempty"`

		// content
		Content *string `json:"content,omitempty"`

		// The time the account was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
