package customer

type Customer struct {
	Name           string
	IsHighPriority bool
}

func NewCustomer(name string, highpriority bool) *Customer {
	return &Customer{name, highpriority}
}

func (c *Customer) GetName() string {
	return c.Name
}
