package platforms

import (
	"github.com/google/uuid"
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
	UserIdentifier uuid.UUID
	Type           PlatformType `json:"type" binding:"required" gorm:"type:INT"`
	Username       string       `json:"username" binding:"required" gorm:"type:VARCHAR(255)"`
	Password       string       `json:"password" binding:"required" gorm:"type:VARCHAR(255)"`
	Nickname       string       `json:"nickname" binding:"required" gorm:"type:VARCHAR(255)"`
	ProfilePicture []byte       `json:"profilepic" binding:"required" gorm:"type:BLOB"`
}
