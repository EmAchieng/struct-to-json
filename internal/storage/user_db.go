package storage

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"structs-to-json-demo/internal/model"
)

// UserDBStore implements UserStore using PostgreSQL via sqlx.
type UserDBStore struct {
	DB *sqlx.DB
}

func NewUserDBStore(db *sqlx.DB) *UserDBStore {
	return &UserDBStore{DB: db}
}

func (s *UserDBStore) CreateUser(user model.User) (model.User, error) {
	query := `INSERT INTO users (username, email, active, created_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := s.DB.QueryRowx(query, user.Username, user.Email, user.Active, time.Now().UTC()).Scan(&user.ID, &user.CreatedAt)
	return user, err
}

func (s *UserDBStore) GetUser(id int) (model.User, error) {
	var user model.User
	query := `SELECT id, username, email, active, created_at FROM users WHERE id=$1`
	err := s.DB.Get(&user, query, id)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *UserDBStore) ListUsers() ([]model.User, error) {
	users := []model.User{}
	query := `SELECT id, username, email, active, created_at FROM users ORDER BY id`
	err := s.DB.Select(&users, query)
	return users, err
}

func (s *UserDBStore) UpdateUser(id int, updated model.User) (model.User, error) {
	var user model.User
	query := `UPDATE users SET username=$1, email=$2, active=$3 WHERE id=$4 RETURNING id, username, email, active, created_at`
	err := s.DB.Get(&user, query, updated.Username, updated.Email, updated.Active, id)
	if err != nil {
		return user, errors.New("user not found")
	}
	return user, nil
}

func (s *UserDBStore) DeleteUser(id int) error {
	res, err := s.DB.Exec(`DELETE FROM users WHERE id=$1`, id)
	if err != nil {
		return err
	}
	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("user not found")
	}
	return nil
}