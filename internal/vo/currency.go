package vo

import "fmt"

type Currency string

const (
	BTC Currency = "btc"
	LTC Currency = "ltc"
	USD Currency = "usd"
	EUR Currency = "eur"
)

var allCurrencies = []Currency{BTC, LTC, USD, EUR}

func NewCurrency(v string) (Currency, error) {
	c := Currency(v)
	for _, a := range allCurrencies {
		if a == c {
			return c, nil
		}
	}
	return "", fmt.Errorf("invalid currency: %s", v)
}

func (c Currency) String() string { return string(c) }
func (c Currency) Upper() string {
	// simple
	b := []byte(c)
	for i := range b {
		if b[i] >= 'a' && b[i] <= 'z' {
			b[i] -= 32
		}
	}
	return string(b)
}
