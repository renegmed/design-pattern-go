package pricing

// NasdaqPricing is the strategy implementation to get prices from Nasdaq.
type NasdaqPricing struct {
	stocks []*Stock
}

func NewNasdaqPricing() *NasdaqPricing {
	return &NasdaqPricing{}
}

func (np *NasdaqPricing) AddStock(stock *Stock) {
	np.stocks = append(np.stocks, stock)
}

func (np *NasdaqPricing) Search() error {
	if len(np.stocks) > 0 {
		np.stocks[0].Price = 2344500
		np.stocks[1].Price = 5439990
	}

	return nil
}

func (np *NasdaqPricing) GetName() string {
	return "Nasdaq"
}
