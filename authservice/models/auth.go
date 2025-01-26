package models

type Auth struct {
	ID       uint64
	Password string
	Salt     string
}
