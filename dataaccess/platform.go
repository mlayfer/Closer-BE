package dataaccess

import (
	"Closer/common/platforms"
	uuid "github.com/satori/go.uuid"
)

type Platform interface {
	GetPlatform(userIdentifier uuid.UUID, platformType platforms.PlatformType) (*platforms.Platform, error)
	GetUserPlatforms(userIdentifier uuid.UUID) ([]*platforms.Platform, error)
	InsertUserPlatform(platform *platforms.Platform) (*platforms.Platform, error)
}
