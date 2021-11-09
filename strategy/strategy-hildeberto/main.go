package main

import (
	"fmt"
	"strategy_design_pattern/strategy"
	"strategy_design_pattern/strategy/pricing"
)

// Supported exchanges
const (
	NASDAQ = "NASDAQ"
	NYSE   = "NYSE"
	TSX    = "TSX"
)

func main() {
	stocks := []*pricing.Stock{
		{
			Exchange: NASDAQ,
			Ticker:   "GOOGL",
			Name:     "Alphabet Inc.",
		}, {
			Exchange: NYSE,
			Ticker:   "GE",
			Name:     "General Electric CO",
		}, {
			Exchange: TSX,
			Ticker:   "SHOP",
			Name:     "Shopify Inc.",
		}, {
			Exchange: NASDAQ,
			Ticker:   "AAPL",
			Name:     "Apple Inc.",
		}, {
			Exchange: TSX,
			Ticker:   "BMO",
			Name:     "Bank of Montreal",
		}, {
			Exchange: TSX,
			Ticker:   "CP",
			Name:     "Canadian Pacific Raiway Ltd.",
		},
	}

	stocks = getPrices(stocks)

	for _, stock := range stocks {
		fmt.Printf("%s: %d\n", stock.Ticker, stock.Price)
	}
}

func getPrices(stocks []*pricing.Stock) []*pricing.Stock {
	exchangeFactory := prepareExchangeStrategies(stocks)
	return searchStocks(exchangeFactory, stocks)
}

func prepareExchangeStrategies(stocks []*pricing.Stock) *strategy.ExchangeFactory {
	exchangeFactory := &strategy.ExchangeFactory{}
	for _, stock := range stocks {
		pricing := exchangeFactory.GetExchangePricing(stock)
		pricing.AddStock(stock)
	}
	return exchangeFactory
}

func searchStocks(exchangeFactory *strategy.ExchangeFactory, stocks []*pricing.Stock) []*pricing.Stock {
	for _, exchange := range exchangeFactory.GetExchanges() {
		err := exchange.Search()
		if err != nil {
			fmt.Printf("Couldn't fetch prices from %s.\n", exchange.GetName())
		}
	}
	return stocks
}
