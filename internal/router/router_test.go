package router

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestNotFoundRoute(t *testing.T) {
    handler := NewWithDB(nil) // Pass nil or a mock DB if needed

    req := httptest.NewRequest("GET", "/does-not-exist", nil)
    w := httptest.NewRecorder()
    handler.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Errorf("expected 404, got %d", w.Code)
    }
}