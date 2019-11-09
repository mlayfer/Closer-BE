package users

import (
	"Closer/platforms"
	"github.com/google/uuid"
)

type User struct {
	Identifier uuid.UUID
	FirstName  string
	LastName   string
	Email      string
	Platforms  []*platforms.Platform
}

func NewUser(FirstName, LastName, Email string, platforms []*platforms.Platform) *User {
	return &User{
		Identifier: uuid.New(),
		FirstName:  FirstName,
		LastName:   LastName,
		Email:      Email,
		Platforms:  platforms,
	}
}
