package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/user/go-exchange-rate/internal/vo"
)

type RateProvider interface {
	GetRate(c1, c2 vo.Currency) (vo.Number, error)
}

type BitfinexRateProvider struct {
	Client *http.Client
	URL    string
}

func NewBitfinexRateProvider(client *http.Client) *BitfinexRateProvider {
	if client == nil {
		client = http.DefaultClient
	}
	return &BitfinexRateProvider{Client: client, URL: "https://api.bitfinex.com/v2/calc/fx"}
}

func (p *BitfinexRateProvider) GetRate(c1, c2 vo.Currency) (vo.Number, error) {
	body, _ := json.Marshal(map[string]string{
		"ccy1": c1.Upper(),
		"ccy2": c2.Upper(),
	})
	resp, err := p.Client.Post(p.URL, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var arr []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&arr); err != nil {
		return "", err
	}
	if len(arr) == 0 {
		return "", fmt.Errorf("empty response")
	}
	rateStr := fmt.Sprintf("%v", arr[0])
	return vo.NewNumber(rateStr)
}
