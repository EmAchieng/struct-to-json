package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"structs-to-json-demo/internal/model"
	"structs-to-json-demo/internal/storage"
)

// UserHandler contains user logic.
type UserHandler struct {
	Store storage.UserStore
}

func NewUserHandler(store storage.UserStore) *UserHandler {
	return &UserHandler{Store: store}
}

func (h *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Store.ListUsers()
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Could not list users")
		return
	}
	RespondJSON(w, http.StatusOK, users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	user, err := h.Store.GetUser(id)
	if err != nil {
		RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	if err := user.Validate(); err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	created, err := h.Store.CreateUser(user)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "Could not create user")
		return
	}
	RespondJSON(w, http.StatusCreated, created)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var update model.User
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}
	if err := update.Validate(); err != nil {
		RespondError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	user, err := h.Store.UpdateUser(id, update)
	if err != nil {
		RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	RespondJSON(w, http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if err := h.Store.DeleteUser(id); err != nil {
		RespondError(w, http.StatusNotFound, "User not found")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// --- Helpers ---

func RespondJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func RespondError(w http.ResponseWriter, code int, message string) {
	RespondJSON(w, code, map[string]string{"error": message})
}

// extractID parses /users/{id}
func extractID(path string) (int, error) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 2 {
		return 0, errors.New("no id")
	}
	return strconv.Atoi(parts[len(parts)-1])
}