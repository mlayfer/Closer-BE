package dataaccess

import (
	"Closer/common/users"
	uuid "github.com/satori/go.uuid"
)

type User interface {
	GetAllUsers() ([]*users.User, error)
	GetByIdentifier(identifier uuid.UUID) (*users.User, error)
	InsertUser(user *users.User) error
	DeleteUser(identifier uuid.UUID) error
}
