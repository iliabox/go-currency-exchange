package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	h := NewRateHandler(nil)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ping", nil)
	h.Ping(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatal(rr.Code)
	}
}

func TestIndex(t *testing.T) {
	h := NewRateHandler(nil)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	h.Index(rr, req)
	if rr.Code != http.StatusOK {
		t.Fatal(rr.Code)
	}
}
