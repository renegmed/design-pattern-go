package main

// const (
// 	pepperoniPrice = 1.75
// 	artichokePrice = 2.00
// 	anchoviePrice = 2.50
// 	tomatoePrice = 0.75

// 	pepperoniCalories = 200
// 	artichokeCalories = 150
// 	anchovieCalories = 350
// 	tomatoeCalories = 50
// )

// var (
// 	pepperoniPrice float64
// 	artichokePrice float64
// 	anchoviePrice  float64
// 	tomatoePrice   float64

// 	pepperoniCalories int
// 	artichokeCalories int
// 	anchovieCalories  int
// 	tomatoeCalories   int
// )

type pepperoniTopping struct {
	pizza        pizza
	toppingprice float64
	calories     int
}
type artichokeTopping struct {
	pizza        pizza
	toppingprice float64
	calories     int
}
type anchovieTopping struct {
	pizza        pizza
	toppingprice float64
	calories     int
}
type tomatoeTopping struct {
	pizza        pizza
	toppingprice float64
	calories     int
}

func (p *pepperoniTopping) getPrice() float64 {
	return p.pizza.getPrice() + p.toppingprice
}
func (a *artichokeTopping) getPrice() float64 {
	return a.pizza.getPrice() + a.toppingprice
}
func (a *anchovieTopping) getPrice() float64 {
	return a.pizza.getPrice() + a.toppingprice
}
func (t *tomatoeTopping) getPrice() float64 {
	return t.pizza.getPrice() + t.toppingprice
}

func (p *pepperoniTopping) getCalories() int {
	return p.pizza.getCalories() + p.calories
}
func (a *artichokeTopping) getCalories() int {
	return a.pizza.getCalories() + a.calories
}
func (a *anchovieTopping) getCalories() int {
	return a.pizza.getCalories() + a.calories
}
func (t *tomatoeTopping) getCalories() int {
	return t.pizza.getCalories() + t.calories
}
