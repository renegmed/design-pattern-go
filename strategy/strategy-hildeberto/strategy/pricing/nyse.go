package pricing

// NYSEPricing is the strategy implementation to get prices from NYSE.
type NYSEPricing struct {
	stocks []*Stock
}

func NewNYSEPricing() *NYSEPricing {
	return &NYSEPricing{}
}

func (nyp *NYSEPricing) AddStock(stock *Stock) {
	nyp.stocks = append(nyp.stocks, stock)
}

func (nyp *NYSEPricing) Search() error {
	if len(nyp.stocks) > 0 {
		nyp.stocks[0].Price = 344500
	}

	return nil
}

func (nyp *NYSEPricing) GetName() string {
	return "New York Stock Exchange"
}
