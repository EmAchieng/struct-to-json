package model

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

type User struct {
	ID        int       `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Active    bool      `db:"active" json:"active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (u *User) Validate() error {
	if len(strings.TrimSpace(u.Username)) < 3 {
		return errors.New("username must be at least 3 characters")
	}
	if !isValidEmail(u.Email) {
		return errors.New("invalid email format")
	}
	return nil
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

func isValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func (u *User) UpdateFrom(in User) {
	if in.Username != "" {
		u.Username = in.Username
	}
	if in.Email != "" {
		u.Email = in.Email
	}
	u.Active = in.Active
}