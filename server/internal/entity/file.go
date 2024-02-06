package entity

// Binary entity of binary data
type Binary struct {
	UserID string `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title" validate:"required,lte=30"`
	Data   []byte `json:"data" db:"data" validate:"omitempty"`
}

