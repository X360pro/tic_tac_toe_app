package service

import (
	"components"
)

type BoardService struct {
	*components.Board
}

func NewBoardService() *BoardService {
	return &BoardService{components.NewBoard(3)}
}

func (b *BoardService) PutMarkInPosition(position uint8, player components.Player) error {
	err := b.Cells[position].SetMark(player.Mark)
	if err != nil {
		return err
	}
	return nil
}

func (b *BoardService) CheckBoardIsFull() bool {
	for i := 0; i < int(b.Size*b.Size); i++ {
		if b.Cells[i].GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
func (b *BoardService) PrintBoard() string {
	out := ""
	for i := 0; i < int(b.Size*b.Size); i++ {
		out += b.Cells[i].Mark
		if i%int(b.Size) == int(b.Size)-1 {
			out += "\n"
		}
	}
	return out
}
