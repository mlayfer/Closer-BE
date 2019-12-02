package users

import (
	"Closer/common/platforms"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
	Identifier uuid.UUID             `json:"identifier" binding:"required" gorm:"type:uuid;primary_key;"`
	FirstName  string                `json:"firstname" binding:"required" gorm:"type:VARCHAR(255)"`
	LastName   string                `json:"lastname" binding:"required" gorm:"type:VARCHAR(255)"`
	Email      string                `json:"email" binding:"required" gorm:"VARCHAR(255)"`
	Platforms  []*platforms.Platform `json:"platforms" binding:"required" gorm:"foreignkey:UserIdentifier"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	userID := uuid.New()
	return scope.SetColumn("ID", userID)
}
