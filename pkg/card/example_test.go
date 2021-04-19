package card

import (
	"fmt"

	"github.com/Hikmatullo1/app/pkg/types")


func WithdrawExampe() {
	card := types.Card{
		ID: 1,
		Name: "aaaaa",
		Active: true,
		Balance: 10_000,
		Category: "computer",
		Currency: "TJS",
	}
	Withdraw(&card, 10_000)
	fmt.Println(card)

	// Output: 20_000
}