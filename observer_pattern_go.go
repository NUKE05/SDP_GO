package main

import "fmt"

type Observer interface {
	Update(balance float64)
}

type Client struct {
	name    string
	balance float64
}

func (c *Client) Update(balance float64) {
	c.balance = balance
	fmt.Printf("Client %s: Updated balance: $%.2f\n", c.name, c.balance)
}

type Bank struct {
	observers []Observer
	balance   float64
}

func (b *Bank) RegisterObserver(observer Observer) {
	b.observers = append(b.observers, observer)
}

func (b *Bank) RemoveObserver(observer Observer) {
	for i, o := range b.observers {
		if o == observer {
			b.observers = append(b.observers[:i], b.observers[i+1:]...)
			break
		}
	}
}

func (b *Bank) SetBalance(balance float64) {
	b.balance = balance
	b.notifyObservers()
}

func (b *Bank) notifyObservers() {
	for _, observer := range b.observers {
		observer.Update(b.balance)
	}
}

func main() {
	bank := &Bank{}
	client1 := &Client{name: "Alice"}
	client2 := &Client{name: "Bob"}

	bank.RegisterObserver(client1)
	bank.RegisterObserver(client2)

	bank.SetBalance(1000.00)
	bank.SetBalance(1500.50)

	bank.RemoveObserver(client1)

	bank.SetBalance(2000.75)
}
