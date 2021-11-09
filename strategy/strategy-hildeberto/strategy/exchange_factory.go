package strategy

import (
	"strategy_design_pattern/strategy/pricing"
)

// ExchangeFactory holds the pricing strategies for the duration of the execution.
type ExchangeFactory struct {
	exchanges map[string]pricing.Pricing
}

// GetExchangePricing returns the Pricing based on the stock exchange.
func (ec *ExchangeFactory) GetExchangePricing(stock *pricing.Stock) pricing.Pricing {
	if ec.exchanges == nil {
		ec.exchanges = make(map[string]pricing.Pricing)
		ec.exchanges["NASDAQ"] = pricing.NewNasdaqPricing() //new(NasdaqPricing)
		ec.exchanges["NYSE"] = pricing.NewNYSEPricing()     //new(NYSEPricing)
		ec.exchanges["TSX"] = pricing.NewTsxPricing()       //new(TsxPricing)
	}
	return ec.exchanges[stock.Exchange]
}

func (ec *ExchangeFactory) GetExchanges() map[string]pricing.Pricing {
	return ec.exchanges
}
