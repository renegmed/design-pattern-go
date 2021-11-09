package menu

type Sub struct {
	bread     string
	hasCheese bool
	toppings  []string
	sauces    []string
}

func (s *Sub) SetBread(bread string) {
	s.bread = bread //"italian"
}

func (s *Sub) GetBread() string {
	return s.bread
}
func (s *Sub) SetCheese(hasCheese bool) {
	s.hasCheese = hasCheese
}

func (s *Sub) HasCheese() bool {
	return s.hasCheese
}
func (s *Sub) SetToppings(toppings []string) {
	s.toppings = toppings //[]string{"roasted chicken", "olives", "onions", "jalapeños"}
}

func (s *Sub) AddTopping(topping string) {
	s.toppings = append(s.toppings, topping)
}
func (s *Sub) GetToppings() []string {
	return s.toppings
}

func (s *Sub) SetSauces(sauces []string) {
	s.sauces = sauces // []string{"chilli", "bbq"}
}

func (s *Sub) AddSauce(sauce string) {
	s.sauces = append(s.sauces, sauce)
}
func (s *Sub) GetSauces() []string {
	return s.sauces
}
func (s *Sub) GetSub() Sub {
	return *s
}
