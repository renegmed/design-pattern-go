package main

// The restaurant contains the total dishes and the total cleaned dishes
type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

// `NewRestaurant` constructs a new restaurant instance with 10 dishes,
// all of them being clean
func NewRestaurant(totalDishes, totalCleanDishes int) *Restaurant {
	//const totalDishes = 10
	return &Restaurant{
		TotalDishes:   totalDishes,
		CleanedDishes: totalCleanDishes,
	}
}
func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		restaurant: r,
		n:          n,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{
		restaurant: r,
	}
}
