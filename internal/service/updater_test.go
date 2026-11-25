package service

import (
	"context"
	"testing"
	"time"

	"github.com/user/go-exchange-rate/internal/domain"
	"github.com/user/go-exchange-rate/internal/provider"
	"github.com/user/go-exchange-rate/internal/repository"
	"github.com/user/go-exchange-rate/internal/vo"
)

type fakeProvider struct{ rate vo.Number }

func (f fakeProvider) GetRate(a, b vo.Currency) (vo.Number, error) { return f.rate, nil }

type fakeRepo struct{ saved int }

func (f *fakeRepo) Save(ctx context.Context, e *domain.CurrentExchangeRate) error {
	f.saved++
	return nil
}
func (f *fakeRepo) Find(ctx context.Context, c1, c2 vo.Currency, from, to *time.Time) ([]domain.CurrentExchangeRate, error) {
	return nil, nil
}

var _ provider.RateProvider = fakeProvider{}
var _ repository.ExchangeRateRepository = &fakeRepo{}

func TestUpdaterRun(t *testing.T) {
	n, _ := vo.NewNumber("1.5")
	p := fakeProvider{rate: n}
	r := &fakeRepo{}
	u := NewExchangeRateUpdater(p, r, [][2]vo.Currency{{vo.BTC, vo.USD}})
	if err := u.Run(context.Background()); err != nil {
		t.Fatal(err)
	}
	if r.saved != 1 {
		t.Fatalf("expected 1 saved, got %d", r.saved)
	}
}
