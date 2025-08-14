package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (c CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("purchase is valid ,Paid %f using credit card with number %s", amount, c.CardNumber)
}

type Cash struct{}

func (Cash) Pay(amount float64) string {
	return fmt.Sprintf("purchase is valid , Paid %f using cash", amount)
}

func main() {

	card := CreditCard{CardNumber: "5022291045867522"}
	cash := Cash{}

	paymentMethod := []PaymentMethod{card, cash}
	for _, method := range paymentMethod {
		fmt.Println(method.Pay(100000))
	}
}
