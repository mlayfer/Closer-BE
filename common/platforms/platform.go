package platforms

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PlatformType int

const (
	_                     = iota // ignore 0
	Facebook PlatformType = iota
	Instagram
	Linkedin
	Snapchat
	Tiktok
	Youtube
	Pinterest
	Google
)

type Platform struct {
	Identifier	   uuid.UUID 	`json:"-" gorm:"primary_key;"`
	UserIdentifier uuid.UUID	`json:"-"`
	Type           PlatformType `json:"type" binding:"required"`
	Username       string       `json:"username" binding:"required"`
	Password       string       `json:"password" binding:"required"`
	Nickname       string       `json:"nickname" binding:"required"`
	ProfilePicture []byte       `json:"profilepic" binding:"required"`
}

func (p *Platform) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("Identifier", uuid.NewV4().String())
}
