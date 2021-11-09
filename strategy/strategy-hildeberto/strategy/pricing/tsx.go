package pricing

// TsxPricing is the strategy implementation to get prices from Tsx.
type TsxPricing struct {
	stocks []*Stock
}

func NewTsxPricing() *TsxPricing {
	return &TsxPricing{}
}

func (tp *TsxPricing) AddStock(stock *Stock) {
	tp.stocks = append(tp.stocks, stock)
}

func (tp *TsxPricing) Search() error {
	if len(tp.stocks) > 0 {
		tp.stocks[0].Price = 8344500
		tp.stocks[1].Price = 239990
		tp.stocks[2].Price = 39990
	}

	return nil
}

func (tp *TsxPricing) GetName() string {
	return "Toronto Stock Exchange"
}
