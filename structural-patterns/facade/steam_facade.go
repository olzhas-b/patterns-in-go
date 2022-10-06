package facade

import "fmt"

// facade is a structural pattern
// which provide simple interface to client
// and interact with multiple classes or library
// to achieve the goal
type SteamFacade struct {
	Steam
	payMethod IPay
	Notification
}

func NewSteamFacade(payMethod IPay) *SteamFacade {
	return &SteamFacade{
		Steam: Steam{
			Account: Account{},
		},
		payMethod: &CreditCard{
			money: 10000,
		},
		Notification: Notification{
			email: "oljas.bazarbekov15@gmail.com",
		},
	}
}
func (stFacade *SteamFacade) BuyGame(game Game) error {
	if stFacade.Account.checkAlreadyHaveGame(game) {
		return fmt.Errorf("client already have game: %v", game)
	}

	if stFacade.payMethod.Pay(game.Cost) {
		stFacade.Account.addGame(game)
	}

	msg := fmt.Sprintf("You succeffully bought %s", game)
	if err := stFacade.notifyByEmail(msg); err != nil {
		return fmt.Errorf("notifyByEmail got error %v", err)
	}
	return nil
}
