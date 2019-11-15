package users

import (
	"Closer/common/platforms"
	"github.com/google/uuid"
)

type User struct {
	Identifier uuid.UUID             `json:"identifier" binding:"required"`
	FirstName  string                `json:"username" binding:"required"`
	LastName   string                `json:"lastname" binding:"required"`
	Email      string                `json:"email" binding:"required"`
	Platforms  []*platforms.Platform `json:"platforms" binding:"required"`
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
