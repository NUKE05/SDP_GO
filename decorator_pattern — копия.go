package main

import "fmt"

// Интерфейс Кофе определяющий методы для получения цены и описания
type Coffee interface {
	Cost() int
	Description() string
}

// BaseCoffee - обычный кофе
type BaseCoffee struct{}

func (c BaseCoffee) Cost() int {
	return 790 // стоимость обычного кофе
}

func (c BaseCoffee) Description() string {
	return "Кофе"
}

// Декоратор добавляющий молоко
type MilkDecorator struct {
	Coffee Coffee
}

func (m MilkDecorator) Cost() int {
	return m.Coffee.Cost() + 100 // Добавление стоимости молока
}

func (m MilkDecorator) Description() string {
	return m.Coffee.Description() + ", Молоко"
}

// Декоратор добавляющий сахар
type SugarDecorator struct {
	Coffee Coffee
}

func (s SugarDecorator) Cost() int {
	return s.Coffee.Cost() + 20 // Добавление стоимости сахар
}

func (s SugarDecorator) Description() string {
	return s.Coffee.Description() + ", Сахар"
}

// Добавление сиропа
type Syrup struct {
	Coffee Coffee
}

func (sy Syrup) Cost() int {
	return sy.Coffee.Cost() + 50 // Добавление стоимости сиропа
}

func (sy Syrup) Description() string {
	return sy.Coffee.Description() + ", Сироп"
}

func main() {
	coffee := BaseCoffee{}
	fmt.Printf("Стомость: %d Тенге\n", coffee.Cost())
	fmt.Printf("Состав: %s\n", coffee.Description())

	coffeeWithMilk := MilkDecorator{Coffee: coffee}
	fmt.Printf("\nСтоимость: %d Тенге\n", coffeeWithMilk.Cost())
	fmt.Printf("Состав: %s\n", coffeeWithMilk.Description())

	coffeeWithMilkSugarVanilla := Syrup{Coffee: SugarDecorator{Coffee: coffeeWithMilk}}
	fmt.Printf("\nСтоимость: %d Тенге\n", coffeeWithMilkSugarVanilla.Cost())
	fmt.Printf("Description: %s\n", coffeeWithMilkSugarVanilla.Description())
}
