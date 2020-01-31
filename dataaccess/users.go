package dataaccess

import (
	"Closer/common/users"
	"fmt"
	"github.com/google/uuid"
)

func (db *Database) GetAllUsers() ([]*users.User, error) {
	panic("implement me")
}

func (db *Database) DeleteUser(id uuid.UUID) error{
	panic("implement me")
}

func (db *Database) GetUserByID(id uuid.UUID) (*users.User, error){
	panic("implement me")
}

func (db *Database) InsertUser(user *users.User) error{
	if ok := db.db.NewRecord(user); ok{
		return fmt.Errorf("userID %s is already in use", user.Identifier)
	}

	if err := db.db.Create(&user).Error; err != nil {
		return fmt.Errorf("Unable to create user. err: %s", err)
	}

	return nil
}