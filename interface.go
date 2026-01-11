package main

import "fmt"

type Card struct{}
type Cash struct{}

type Payment interface {
	Pay(amount int)
}

func (card *Card) Pay(amount int) {
	fmt.Printf("Оплата картой: %d \n", amount)
}

func (card *Cash) Pay(amount int) {
	fmt.Printf("Оплата наличными: %d \n", amount)
}

func notMain() {
	payByCard := Card{}
	payByCash := Cash{}

	Process(&payByCard, 100)
	Process(&payByCash, 200)
}

func Process(p Payment, amount int) {
	p.Pay(amount)
}
