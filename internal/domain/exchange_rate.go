package domain

import (
	"time"

	"github.com/google/uuid"

	"github.com/user/go-exchange-rate/internal/vo"
)

type CurrentExchangeRate struct {
	ID        uuid.UUID
	Currency1 vo.Currency
	Currency2 vo.Currency
	Rate      vo.Number
	CreatedAt time.Time
}

func NewCurrentExchangeRate(c1, c2 vo.Currency, rate vo.Number) *CurrentExchangeRate {
	return &CurrentExchangeRate{
		ID:        uuid.New(),
		Currency1: c1,
		Currency2: c2,
		Rate:      rate,
		CreatedAt: time.Now().UTC(),
	}
}
