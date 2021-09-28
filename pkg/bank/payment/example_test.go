package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources()  {
	
	cards:= []types.Card {
		{
			PAN: "1233 3123 4344 3443",
			Balance: 300,
			Name: "Visa",
			Active: true,
		},
		{
			PAN: "1231 3244 6756 5675",
			Balance: 400,
			Name: "Master",
			Active: true,
		},
	}

	paymentSources:=PaymentSources(cards)

	for i := 0; i < len(paymentSources); i++ {
		fmt.Println(paymentSources[i].Number)
	}

	
	//Output:
	//1233 3123 4344 3443
	//1231 3244 6756 5675
}