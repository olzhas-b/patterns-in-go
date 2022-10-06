package facade

import "fmt"

type Game struct {
	ID   int64
	Name string
	Cost float64
}

func (game Game) String() string {
	return fmt.Sprintf("game name is %s, ID%d, cost: %f", game.Name, game.ID, game.Cost)
}
