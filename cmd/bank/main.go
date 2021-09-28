package main

import (
	"bank/pkg/bank/payment"
	"bank/pkg/bank/types"
	"fmt"
)

func main() {
	cards := []types.Card{
		{
			PAN:     "1233 3123 4344 3443",
			Balance: 2,
			Name:    "Visa",
			Active:  true,
		},
		{
			PAN:     "1231 3244 6756 5675",
			Balance: 400,
			Name:    "Master",
			Active:  true,
		},
		{
			PAN:     "1234 2342 1211 1111",
			Balance: 100,
			Name:    "UnionPay",
			Active:  true,
		},
	}
	paymentSources := payment.PaymentSources(cards)

	for i := 0; i < len(paymentSources); i++ {
		fmt.Println(paymentSources[i].Number)
	}

	// fmt.Println(len(cards))
	// fmt.Println(payment.PaymentSources(cards))

}
