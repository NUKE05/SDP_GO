package main

import "fmt"

type Beshparmak interface {
	Display() string
}

type Sorpa interface {
	Display() string
}

type Pizza struct{}

func (p Pizza) Display() string {
	return "Pizza"
}

type Pasta struct{}

func (p Pasta) Display() string {
	return "Pasta"
}

type Cola struct{}

func (s Cola) Display() string {
	return "Soft Drink"
}

type Wine struct{}

func (w Wine) Display() string {
	return "Wine"
}

type MenuFactory interface {
	CreateBeshmarmak() Beshparmak
	CreateSorpa() Sorpa
}

type ItalianMenuFactory struct{}

func (f ItalianMenuFactory) CreateBeshmarmak() Beshparmak {
	return Pasta{}
}

func (f ItalianMenuFactory) CreateSorpa() Sorpa {
	return Wine{}
}

type FastFoodMenuFactory struct{}

func (f FastFoodMenuFactory) CreateBeshmarmak() Beshparmak {
	return Pizza{}
}

func (f FastFoodMenuFactory) CreateSorpa() Sorpa {
	return Cola{}
}

func createOrder(factory MenuFactory) {
	beshmarmak := factory.CreateBeshmarmak()
	Sorpa := factory.CreateSorpa()
	fmt.Printf("Order: %s with %s\n", beshmarmak.Display(), Sorpa.Display())
}

func main() {
	italianFactory := ItalianMenuFactory{}
	createOrder(italianFactory)

	fastFoodFactory := FastFoodMenuFactory{}
	createOrder(fastFoodFactory)
}
