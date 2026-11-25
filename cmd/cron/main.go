package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/user/go-exchange-rate/internal/config"
	"github.com/user/go-exchange-rate/internal/provider"
	"github.com/user/go-exchange-rate/internal/repository"
	"github.com/user/go-exchange-rate/internal/service"
	"github.com/user/go-exchange-rate/internal/vo"
)

func main() {
	cfg := config.Load()
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	repo := repository.NewPostgresRepository(pool)
	prov := provider.NewBitfinexRateProvider(nil)
	pairs := [][2]vo.Currency{
		{vo.BTC, vo.LTC},
		{vo.LTC, vo.BTC},
		{vo.BTC, vo.USD},
		{vo.USD, vo.BTC},
		{vo.BTC, vo.EUR},
		{vo.EUR, vo.BTC},
	}
	updater := service.NewExchangeRateUpdater(prov, repo, pairs)
	if err := updater.Run(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
}
