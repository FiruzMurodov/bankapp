package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleTotal() {
	card:= []types.Card{
		{
			Balance: 10,
			Active: false,
		},
		{
			Balance: -20,
			Active: true,
		},
		{
			Balance: 30,
			Active: true,
		},
	}
fmt.Println(Total(card))
	//Output:
	//30
}