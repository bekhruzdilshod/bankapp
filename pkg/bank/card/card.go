package card

import "bank/pkg/bank/types"


// Функция выпуска новой карты 
func IssueCard(currency types.Currency, color string, name string) types.Card {

	card := types.Card {
		ID: 1000,
		PAN: "5058 xxxx xxxx 0001",
		Balance: 0,
		Currency: currency,
		Color: color,
		Name: name,
		Active: true,
	}

	return card

}


// Функция снятия средств со счета карты
func Withdraw(card types.Card, amount types.Money) types.Card {

	if card.Active == true && card.Balance > amount {  // Если карта активная и ее баланс больше 0
		
		if amount > 0 && amount < 2_000_000 {  // Если сумма снятия больше 0 и меньше 20_000
			card.Balance = card.Balance - amount
		}  

	}

	return card
}


// Функция пополнения средств на счете
func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= 5_000_000 && amount > 0 {
		card.Balance += amount
	}
}


// Добавление бонуса в размере 3% годовых за минимальный баланс
func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	bonus := int(card.MinBalance) * percent / 100 * daysInMonth / daysInYear

	if card.Active && card.Balance > 0 && bonus <= 500_000 {
		card.Balance += types.Money(bonus)
	}

}


// Фукнция высчитывает общую сумму операций по переданным в аргументы картам
func Total(cards []types.Card) types.Money {

	sum := types.Money(0)

	for _, card := range cards {
		if card.Balance > 0 && card.Active {
			sum += card.Balance
		}  
	}

	return sum
}