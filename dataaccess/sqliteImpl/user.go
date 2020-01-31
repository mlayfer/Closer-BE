package sqliteImpl

import (
	"Closer/common/users"
	"fmt"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{db: db}
}

func (u *UserDB) GetAllUsers() (users []*users.User, err error) {
	u.db.Find(users)
	err = nil
	return
}

func (u *UserDB) GetByIdentifier(identifier uuid.UUID) (user *users.User, err error) {
	u.db.Where(&users.User{	Identifier: identifier }).First(user)
	err = nil
	return
}

func (u *UserDB) InsertUser(user *users.User) error {
	if !u.db.NewRecord(user) {
		return fmt.Errorf("cannot insert user, user already exists")
	}

	u.db.Create(user)
	return nil
}
