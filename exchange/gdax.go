package exchange

import (
	"github.com/preichenberger/go-gdax"
	"fmt"
)

type GDAX struct {
	Exchange
}

func (exchange GDAX) GetPrice(base, quote string) (*Response, *Error) {
	clientInterface := exchange.GetConfig().Client
	client := clientInterface.(*gdax.Client)

	ticker, err := client.GetTicker(fmt.Sprintf("%s-%s", base, quote))
	if err != nil {
		return nil, &Error{exchange.GetConfig().Name, "500 ERROR", err.Error()}
	}

	return &Response{exchange.GetConfig().Name, ticker.Price,  ticker.Volume}, nil
}

func (exchange GDAX) GetPairs() []*Pair {
	clientInterface := exchange.GetConfig().Client
	client := clientInterface.(*gdax.Client)

	products, err := client.GetProducts()
	if err != nil {
		return []*Pair{}
	}
	var pairs []*Pair
	for _, product := range products {
		pairs = append(pairs, &Pair{product.BaseCurrency, product.QuoteCurrency})
	}
	return pairs
}

func (exchange GDAX) GetConfig() *Config {
	return &Config{Name: "GDAX", Client: gdax.NewClient("", "", "")}
}