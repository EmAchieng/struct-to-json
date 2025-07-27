package test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "structs-to-json-demo/internal/router"
)

func TestHealthzEndpoint(t *testing.T) {
    // You can pass nil for DB if your router allows it, or mock as needed
    handler := router.NewWithDB(nil)

    req := httptest.NewRequest("GET", "/healthz", nil)
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", w.Code)
    }
    if w.Body.String() != "ok" {
        t.Fatalf("expected body 'ok', got '%s'", w.Body.String())
    }
}