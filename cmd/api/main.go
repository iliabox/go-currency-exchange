package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/user/go-exchange-rate/internal/config"
	"github.com/user/go-exchange-rate/internal/handler"
	"github.com/user/go-exchange-rate/internal/repository"
)

func main() {
	cfg := config.Load()
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := repository.NewPostgresRepository(pool)
	h := handler.NewRateHandler(repo)

	r := chi.NewRouter()
	r.Get("/", h.Index)
	r.Get("/ping", h.Ping)
	r.Get("/rate/{currency1}/{currency2}", h.Rate)
	log.Fatal(http.ListenAndServe(cfg.HTTPAddr, r))
}
