package ether

import (
	"encoding/json"
	"math/big"
	"net/http"
	"time"

	"github.com/tzero-dev/go-t0ken/units"
)

const exchange = "Coin Market Cap"
const exchangeURL = "https://api.coinmarketcap.com/v1/ticker/ethereum/"

type exchangeRate struct {
	USD string `json:"price_usd"`
	BTC string `json:"price_btc"`
}

func getExchangePrice(ether *big.Float) (*big.Float, *big.Float, *big.Float, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	r, err := client.Get(exchangeURL)
	if err != nil {
		return nil, nil, nil, err
	}
	defer r.Body.Close()
	var s []exchangeRate
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		return nil, nil, nil, err
	}

	usd, _, err := new(big.Float).Parse(s[0].USD, 10)
	if err != nil {
		return nil, nil, nil, err
	}
	btc, _, err := new(big.Float).Parse(s[0].BTC, 10)
	if err != nil {
		return nil, nil, nil, err
	}
	usd.Mul(usd, ether)
	btc.Mul(btc, ether)
	sat := units.ConvertFloat(btc, units.Ether, units.Unit(10))
	return usd, btc, sat, err
}
