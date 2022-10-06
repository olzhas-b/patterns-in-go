package facade

import "sync"

type CreditCard struct {
	money         float64
	accountNumber string
	accountName   string
	expireDate    string
	PIN_CODE      string
	mu            sync.Mutex
}

func NewCreditCard() *CreditCard {
	return &CreditCard{
		money: 100000,
	}
}

func (creditCard *CreditCard) Pay(cost float64) bool {
	creditCard.mu.Lock()
	defer creditCard.mu.Unlock()

	if creditCard.money < cost {
		return false
	}
	creditCard.money -= cost
	return true
}
