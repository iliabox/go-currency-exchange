package service

import (
	"context"

	"github.com/user/go-exchange-rate/internal/domain"
	"github.com/user/go-exchange-rate/internal/provider"
	"github.com/user/go-exchange-rate/internal/repository"
	"github.com/user/go-exchange-rate/internal/vo"
)

type ExchangeRateUpdater struct {
	Provider provider.RateProvider
	Repo     repository.ExchangeRateRepository
	Pairs    [][2]vo.Currency
}

func NewExchangeRateUpdater(p provider.RateProvider, r repository.ExchangeRateRepository, pairs [][2]vo.Currency) *ExchangeRateUpdater {
	return &ExchangeRateUpdater{Provider: p, Repo: r, Pairs: pairs}
}

func (u *ExchangeRateUpdater) Run(ctx context.Context) error {
	for _, pair := range u.Pairs {
		rate, err := u.Provider.GetRate(pair[0], pair[1])
		if err != nil {
			return err
		}
		entity := domain.NewCurrentExchangeRate(pair[0], pair[1], rate)
		if err := u.Repo.Save(ctx, entity); err != nil {
			return err
		}
	}
	return nil
}
