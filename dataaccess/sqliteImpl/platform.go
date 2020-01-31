package sqliteImpl

import (
	"Closer/common/platforms"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type PlatformDB struct {
	db *gorm.DB
}

func NewPlatformDB(db *gorm.DB) *PlatformDB {
	return &PlatformDB{db: db}
}

func (p *PlatformDB) GetPlatform(userIdentifier uuid.UUID, platformType platforms.PlatformType) (platform *platforms.Platform, err error) {
	p.db.Where(&platforms.Platform{UserIdentifier: userIdentifier, Type: platformType}).First(platform)
	err = nil
	return
}

func (p *PlatformDB) GetUserPlatforms(userIdentifier uuid.UUID) (platformsResult []*platforms.Platform, err error) {
	p.db.Where(&platforms.Platform{UserIdentifier: userIdentifier}).Find(platformsResult)
	err = nil
	return
}

func (p *PlatformDB) InsertUserPlatform(platform *platforms.Platform) (*platforms.Platform, error) {
	if !p.db.NewRecord(platform) {
		return nil, fmt.Errorf("cannot insert user, user already exists")
	}

	p.db.Create(platform)
	return platform, nil
}
