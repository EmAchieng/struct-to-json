package model

import (
    "testing"
    "time"
)

func TestUser_Validate_Valid(t *testing.T) {
    u := &User{
        Username:  "gopher",
        Email:     "go@example.com",
        Active:    true,
        CreatedAt: time.Now(),
    }
    if err := u.Validate(); err != nil {
        t.Errorf("expected valid user, got error: %v", err)
    }
}

func TestUser_Validate_InvalidUsername(t *testing.T) {
    u := &User{
        Username: "go",
        Email:    "go@example.com",
    }
    err := u.Validate()
    if err == nil || err.Error() != "username must be at least 3 characters" {
        t.Errorf("expected username error, got: %v", err)
    }
}

func TestUser_Validate_InvalidEmail(t *testing.T) {
    u := &User{
    	Username: "gopher",
        Email:    "invalid-email",
    }
    err := u.Validate()
    if err == nil || err.Error() != "invalid email format" {
        t.Errorf("expected email error, got: %v", err)
    }
}

func TestUser_UpdateFrom(t *testing.T) {
    u := &User{
        Username: "olduser",
        Email:    "old@example.com",
        Active:   false,
    }
    in := User{
        Username: "newuser",
        Email:    "new@example.com",
        Active:   true,
    }
    u.UpdateFrom(in)
    if u.Username != "newuser" || u.Email != "new@example.com" || !u.Active {
        t.Errorf("update failed: got %+v", u)
    }
}