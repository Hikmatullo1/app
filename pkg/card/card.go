package card

import "github.com/Hikmatullo1/app/pkg/types"

func Withdraw(card *types.Card, amount int) {
	const withdrawLimit = 20_000_00
	if amount < 0 {
		return
	}
	if amount > withdrawLimit {
		return
	}
	if !card.Active {
		return
	}
	if card.Balance < amount {
		return
	}
	card.Balance -= amount
}