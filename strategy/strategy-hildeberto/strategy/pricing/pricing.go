package pricing

// Pricing shapes the strategy pattern to be applied to specialized structs.
type Pricing interface {
	AddStock(stock *Stock)
	Search() error
	GetName() string
}
