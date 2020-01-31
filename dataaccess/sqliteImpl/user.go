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

func (u *UserDB) GetAllUsers() ([]*users.User, error) {
	usrs := make([]*users.User, 0)
	u.db.Preload("Platforms").Find(&usrs)
	return usrs, nil
}

func (u *UserDB) GetByIdentifier(identifier uuid.UUID) (*users.User, error) {
	usr := &users.User{}
	u.db.Where(&users.User{	Identifier: identifier }).Preload("Platforms").First(usr)
	return usr, nil
}

func (u *UserDB) InsertUser(user *users.User) error {
	if !u.db.NewRecord(user) {
		return fmt.Errorf("cannot insert user, user already exists")
	}

	u.db.Create(user)
	return nil
}

func (u *UserDB)DeleteUser(identifier uuid.UUID) error{
	u.db.Delete(&users.User{Identifier: identifier}, "Identifier LIKE ?", identifier)
	return nil
}
