package vo

import "testing"

func TestCurrencyValid(t *testing.T) {
	c, err := NewCurrency("btc")
	if err != nil || c != BTC {
		t.Fatal(err)
	}
}

func TestCurrencyInvalid(t *testing.T) {
	if _, err := NewCurrency("xxx"); err == nil {
		t.Fatal("expected error")
	}
}
