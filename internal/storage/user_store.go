package storage

import "structs-to-json-demo/internal/model"

// UserStore defines the interface for user data storage.
type UserStore interface {
	CreateUser(model.User) (model.User, error)
	GetUser(int) (model.User, error)
	ListUsers() ([]model.User, error)
	UpdateUser(int, model.User) (model.User, error)
	DeleteUser(int) error
}