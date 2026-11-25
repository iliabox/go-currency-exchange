package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"github.com/user/go-exchange-rate/internal/repository"
	"github.com/user/go-exchange-rate/internal/vo"
)

type RateHandler struct {
	Repo repository.ExchangeRateRepository
}

func NewRateHandler(repo repository.ExchangeRateRepository) *RateHandler {
	return &RateHandler{Repo: repo}
}

func (h *RateHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("pong")
}

func (h *RateHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("")
}

func (h *RateHandler) Rate(w http.ResponseWriter, r *http.Request) {
	c1, err := vo.NewCurrency(chi.URLParam(r, "currency1"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	c2, err := vo.NewCurrency(chi.URLParam(r, "currency2"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	var from, to *time.Time
	if v := r.URL.Query().Get("from"); v != "" {
		t, err := time.Parse("2006-01-02", v)
		if err == nil {
			from = &t
		}
	}
	if v := r.URL.Query().Get("to"); v != "" {
		t, err := time.Parse("2006-01-02", v)
		if err == nil {
			to = &t
		}
	}
	data, err := h.Repo.Find(r.Context(), c1, c2, from, to)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
