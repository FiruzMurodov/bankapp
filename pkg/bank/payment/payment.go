package payment

import (
	"bank/pkg/bank/types"
	// "fmt"
)

func Max(payments []types.Payment) types.Payment {
	max := payments[0].Amount
	id:=0
	for i := 0; i < len(payments); i++ {
		if max <= payments[i].Amount {
			max = payments[i].Amount
			id=i
		}
	}
	// fmt.Println(max)

	return payments[id]
}

