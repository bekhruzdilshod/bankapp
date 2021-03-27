package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 2_000_000, Active: true}, 1_000_000)
	fmt.Println(result.Balance)

	// Output: 1000000
}


func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 1_000_000, Active: true}, 2_000_000)
	fmt.Println(result.Balance)

	// Output: 1000000
}


func ExampleWithdraw_inactive() {
	result := Withdraw(types.Card{Balance: 2_000_000, Active: false}, 1_000_000)
	fmt.Println(result.Balance)

	// Output: 2000000
}


func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 5_000_000, Active: true}, 3_000_000)
	fmt.Println(result.Balance)

	// Output: 5000000
}


func ExampleDeposit_positive() {
	card := types.Card{
		Active: true,
		Balance: 1_000_000,
	}

	Deposit(&card, 3_000_000)

	fmt.Println(card.Balance)

	// Output: 4000000
}


func ExampleDeposit_negative() {
	card := types.Card{
		Active: true,
		Balance: 1_000_000,
	}

	Deposit(&card, -3_000_000)

	fmt.Println(card.Balance)

	// Output: 1000000
}


func ExampleDeposit_negativeBalance() {
	card := types.Card{
		Active: true,
		Balance: -1_000_000,
	}

	Deposit(&card, 3_000_000)

	fmt.Println(card.Balance)

	// Output: 2000000
}


func ExampleDeposit_inactive() {
	card := types.Card{
		Active: false,
		Balance: 1_000_000,
	}

	Deposit(&card, 3_000_000)

	fmt.Println(card.Balance)

	// Output: 1000000
}

func ExampleDeposit_limitEqual() {
	card := types.Card{
		Active: true,
		Balance: 1_000_000,
	}

	Deposit(&card, 5_000_000)

	fmt.Println(card.Balance)

	// Output: 6000000
}


func ExampleDeposit_limitAbove() {
	card := types.Card{
		Active: true,
		Balance: 1_000_000,
	}

	Deposit(&card, 6_000_000)

	fmt.Println(card.Balance)

	// Output: 1000000
}


func ExampleAddBonus_normal() {
	card := types.Card{
		Balance: 1_000_000,
		Active: true,
		MinBalance: 1_000_000,
	}

	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	// Output: 1002465
}

func ExampleAddBonus_inactiveCard() {
	card := types.Card{
		Balance: 1_000_000,
		Active: false,
		MinBalance: 1_000_000,
	}

	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	// Output: 1000000
}


func ExampleAddBonus_negativeBalance() {
	card := types.Card{
		Balance: -1_000_000,
		Active: false,
		MinBalance: 1_000_000,
	}

	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	// Output: -1000000
}


func ExampleAddBonus_limitAbove() {
	card := types.Card{
		Balance: 1_000_000,
		Active: false,
		MinBalance: 100_000_000,
	}

	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.Balance)

	// Output: 1000000
}


func ExampleTotal() {
	cards := []types.Card{
		{
		Balance: 1000,
		Active: true,
	},
	{
		Balance: 1000,
		Active: true,
	},
	{
		Balance: 1000,
		Active: false,
	},
	}

	sum := Total(cards) 

	fmt.Println(sum)

	// Output: 2000
}


func ExamplePaymentSources() {
	cards := []types.Card{
		{
			PAN: "123 123 123",
			Balance: 1000,
			Active: true,
		},
		{
			PAN: "233 233 233",
			Balance: 5000,
			Active: false,
		},
		{
			PAN: "433 433 433",
			Balance: 0,
			Active: true,
		},
	}

	for _, paymentSource := range PaymentSources(cards) {
		fmt.Println(paymentSource.Number)
	}

	// Output: 123 123 123

}