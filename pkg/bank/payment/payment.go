package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func Max(payments []types.Payment) types.Payment {
	max := payments[0].Amount
	for i := 0; i < len(payments); i++ {
		if max <= payments[i].Amount {
			max = payments[i].Amount
		}
	}
	fmt.Println(max)

	return payments[0]
}

func PaymentSources(cards []types.Card) []types.PaymentSource {

	paymentSources := []types.PaymentSource{}
	for i := 0; i < len(cards); i++ {

		if cards[i].Active && cards[i].Balance > 0 {
			paymentSource:= []types.PaymentSource{
				{
					Type:    cards[i].Name,
					Number:  string(cards[i].PAN),
					Balance: cards[i].Balance,
				},
			}
			paymentSources = append(paymentSources, paymentSource...)
		}
		
	}
	return paymentSources
}
