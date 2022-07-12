package cash_machine

import (
	"fmt"
	"math"
	"strconv"
)

// £50, £20, £10, £5, £2, £1, 50p, 20p, 10p, 5p, 2p, 1p

func BreakIntoChange(amount float64) map[string]int {
	var change = make(map[string]int)

	roundedAmount := math.Round(amount)
	pounds := int(roundedAmount)
	pences := int(math.Round((amount - roundedAmount) * 100))

	moneyValues := []int{50, 20, 10, 5, 2, 1}

	for _, value := range moneyValues {
		if pounds%value != pounds {
			change["£"+strconv.Itoa(value)] = pounds / value
			pounds = pounds - pounds/value*value
			if pounds == 0 {
				break
			}
		}
	}

	for _, value := range moneyValues {
		if pences%value != pences {
			change[strconv.Itoa(value)+"p"] = pences / value
			pences = pences - pences/value*value
			if pences == 0 {
				break
			}
		}
	}
	fmt.Println(change)
	return change
}
