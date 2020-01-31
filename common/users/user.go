package users

import (
	"Closer/common/platforms"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type User struct {
	Identifier uuid.UUID            `json:"identifier" binding:"required" gorm:"primary_key"`
	FirstName  string               `json:"firstname" binding:"required"`
	LastName   string               `json:"lastname" binding:"required"`
	Email      string               `json:"email" binding:"required"`
	Platforms  []platforms.Platform `gorm:"foreignkey:UserIdentifier;association_foreignkey:Identifier" json:"platforms" binding:"required"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("Identifier", uuid.NewV4().String())
}
