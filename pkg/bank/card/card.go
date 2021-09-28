package card

import (
	"bank/pkg/bank/types"
	
)

func Total(cards []types.Card) types.Money {
	sum := int64(0)

	for i := 0; i < len(cards); i++ {
		if !cards[i].Active || cards[i].Balance < 0 {
			continue
		}

		sum += int64(cards[i].Balance)
	}
	// fmt.Println(sum)

	return types.Money(sum)
}
