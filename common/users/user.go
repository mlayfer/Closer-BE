package users

import (
	"Closer/common/platforms"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Identifier uuid.UUID             `json:"identifier" binding:"required" gorm:"column:identifier;type:VARCHAR(255);primary_key"`
	FirstName  string                `json:"firstname" binding:"required" gorm:"column:first_name;type:VARCHAR(255)"`
	LastName   string                `json:"lastname" binding:"required" gorm:"column:last_name;type:VARCHAR(255)"`
	Email      string                `json:"email" binding:"required" gorm:"column:email;type:VARCHAR(255)"`
	Platforms  []*platforms.Platform `json:"platforms" binding:"required" gorm:"foreignkey:UserIdentifier"`
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
