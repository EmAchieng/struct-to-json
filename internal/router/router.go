package router

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"structs-to-json-demo/internal/handler"
	"structs-to-json-demo/internal/storage"
)

func NewWithDB(db *sqlx.DB) http.Handler {
	store := storage.NewUserDBStore(db)
	userHandler := handler.NewUserHandler(store)
	mux := http.NewServeMux()

	// Health check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// Users collection
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.ListUsers(w, r)
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		default:
			handler.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// Users by ID
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)
		case http.MethodPut:
			userHandler.UpdateUser(w, r)
		case http.MethodDelete:
			userHandler.DeleteUser(w, r)
		default:
			handler.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	// 404 fallback
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler.RespondError(w, http.StatusNotFound, "Not found")
	})

	return mux
}