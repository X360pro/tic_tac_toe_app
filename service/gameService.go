package service

import (
	"components"
	"errors"
)

type GameService struct {
	*ResultService
	player1 *components.Player
	player2 *components.Player
}

func NewGameService(pl_1, pl_2 *components.Player, size int) *GameService {
	return &GameService{NewResultService(NewBoardService(size)), pl_1, pl_2}
}

func (g *GameService) Play(pos uint8, pl *components.Player) (error, string) {
	if pos < 0 || pos > g.Size*g.Size-1 {
		return errors.New("Positoin is out of bounds"), "nil"
	}
	err := g.PutMarkInPosition(pos, pl)
	if err != nil {
		return err, "nil"
	}
	return nil, g.GiveResult(pl)
}
