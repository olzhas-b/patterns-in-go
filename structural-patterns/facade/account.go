package facade

type Account struct {
	login    string
	password string
	games    []Game
}

func (acc *Account) checkAlreadyHaveGame(futureGame Game) bool {
	for _, game := range acc.games {
		if game.ID == futureGame.ID {
			return true
		}
	}
	return false
}

func (acc *Account) addGame(game Game) {
	acc.games = append(acc.games, game)
}
