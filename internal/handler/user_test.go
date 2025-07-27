package handler

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "structs-to-json-demo/internal/model"
)

// MockUserStore implements storage.UserStore for testing
type MockUserStore struct {
    ListUsersFunc   func() ([]model.User, error)
    GetUserFunc     func(int) (model.User, error)
    CreateUserFunc  func(model.User) (model.User, error)
    UpdateUserFunc  func(int, model.User) (model.User, error)
    DeleteUserFunc  func(int) error
}

func (m *MockUserStore) ListUsers() ([]model.User, error) {
    return m.ListUsersFunc()
}
func (m *MockUserStore) GetUser(id int) (model.User, error) {
    return m.GetUserFunc(id)
}
func (m *MockUserStore) CreateUser(u model.User) (model.User, error) {
    return m.CreateUserFunc(u)
}
func (m *MockUserStore) UpdateUser(id int, u model.User) (model.User, error) {
    return m.UpdateUserFunc(id, u)
}
func (m *MockUserStore) DeleteUser(id int) error {
    return m.DeleteUserFunc(id)
}

func TestListUsers_Success(t *testing.T) {
    store := &MockUserStore{
        ListUsersFunc: func() ([]model.User, error) {
            return []model.User{{ID: 1, Username: "gopher"}}, nil
        },
    }
    handler := NewUserHandler(store)
    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()

    handler.ListUsers(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", w.Code)
    }
}

func TestListUsers_Error(t *testing.T) {
    store := &MockUserStore{
        ListUsersFunc: func() ([]model.User, error) {
            return nil, errors.New("fail")
        },
    }
    handler := NewUserHandler(store)
    req := httptest.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()

    handler.ListUsers(w, req)

    if w.Code != http.StatusInternalServerError {
        t.Errorf("expected status 500, got %d", w.Code)
    }
}

func TestGetUser_Success(t *testing.T) {
    store := &MockUserStore{
        GetUserFunc: func(id int) (model.User, error) {
            return model.User{ID: id, Username: "gopher"}, nil
        },
    }
    handler := NewUserHandler(store)
    req := httptest.NewRequest("GET", "/users/1", nil)
    w := httptest.NewRecorder()

    handler.GetUser(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("expected status 200, got %d", w.Code)
    }
}

func TestGetUser_InvalidID(t *testing.T) {
    store := &MockUserStore{}
    handler := NewUserHandler(store)
    req := httptest.NewRequest("GET", "/users/", nil)
    w := httptest.NewRecorder()

    handler.GetUser(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected status 400, got %d", w.Code)
    }
}

func TestGetUser_NotFound(t *testing.T) {
    store := &MockUserStore{
        GetUserFunc: func(id int) (model.User, error) {
            return model.User{}, errors.New("not found")
        },
    }
    handler := NewUserHandler(store)
    req := httptest.NewRequest("GET", "/users/99", nil)
    w := httptest.NewRecorder()

    handler.GetUser(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected status 404, got %d", w.Code)
    }
}

func TestCreateUser_Success(t *testing.T) {
    store := &MockUserStore{
        CreateUserFunc: func(u model.User) (model.User, error) {
            return u, nil
        },
    }
    handler := NewUserHandler(store)
    user := model.User{ID: 2, Username: "alice", Email: "alice@example.com", Active: true}
    body, _ := json.Marshal(user)
    req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
    w := httptest.NewRecorder()

    handler.CreateUser(w, req)

    if w.Code != http.StatusCreated {
        t.Errorf("expected status 201, got %d", w.Code)
    }
}

func TestCreateUser_InvalidJSON(t *testing.T) {
    store := &MockUserStore{}
    handler := NewUserHandler(store)
    req := httptest.NewRequest("POST", "/users", bytes.NewReader([]byte("{invalid")))
    w := httptest.NewRecorder()

    handler.CreateUser(w, req)

    if w.Code != http.StatusBadRequest {
        t.Errorf("expected status 400, got %d", w.Code)
    }
}

func TestCreateUser_StoreError(t *testing.T) {
    store := &MockUserStore{
        CreateUserFunc: func(u model.User) (model.User, error) {
            return model.User{}, errors.New("fail")
        },
    }
    handler := NewUserHandler(store)
    user := model.User{ID: 2, Username: "alice", Email: "alice@example.com", Active: true}
    body, _ := json.Marshal(user)
    req := httptest.NewRequest("POST", "/users", bytes.NewReader(body))
    w := httptest.NewRecorder()

    handler.CreateUser(w, req)

    if w.Code != http.StatusInternalServerError {
        t.Errorf("expected status 500, got %d", w.Code)
    }
}