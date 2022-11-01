/*
Adapter allows us to use a general interface, and other types implementing it
with slightly different functionalities can work normally and still take advantage
of the abstraction.
*/

package main

import "fmt"

type IPayment interface {
	Pay()
}

type Cash struct{}

func (Cash) Pay() {
	fmt.Println("payment done with cash")
}

type BankPayment struct{}

func (BankPayment) Pay(accNumber int) {
	fmt.Println("payment done with account", accNumber)
}

type BankPaymentAdapter struct {
	BankPayment
	accNumber int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.accNumber)
}

func Pay(p IPayment) {
	p.Pay()
}

func main() {
	c := &Cash{}
	b := &BankPaymentAdapter{
		accNumber:   5,
		BankPayment: BankPayment{},
	}

	Pay(c)
	Pay(b)
}
