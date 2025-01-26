package models

type UserAuth struct {
	ID        uint64 `json:"userId"`   // Primary Key
	Username  string `json:"username"` // Unique index
	Email     string `json:"email"`    // Unique index
	Password  string // salted hash
	Salt      string
	Role      uint8  `json:"role"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastname"`
	MediaURL  string `json:"mediaUrl"`
}
