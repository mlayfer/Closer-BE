package platforms

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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
	gorm.Model
	UserIdentifier uuid.UUID
	Type           PlatformType `json:"type" binding:"required" gorm:"type:INT"`
	Username       string       `json:"username" binding:"required" gorm:"column:username;type:VARCHAR(255)"`
	Password       string       `json:"password" binding:"required" gorm:"column:password;type:VARCHAR(255)"`
	Nickname       string       `json:"nickname" binding:"required" gorm:"column:nickname;type:VARCHAR(255)"`
	ProfilePicture []byte       `json:"profilepic" binding:"required" gorm:"column:profile_picture;type:BLOB"`
}
