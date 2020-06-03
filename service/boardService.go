package service

import (
	"components"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(size int) *BoardService {
	return &BoardService{components.NewBoard(uint8(size))}
}

func (b *BoardService) PutMarkInPosition(position uint8, mark string) error {
	err := b.Cells[position].SetMark(mark)
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
		out += " " + b.Cells[i].Mark + " "
		if i%int(b.Size) == int(b.Size)-1 {
			out += "\n"
		}
	}
	return out
}
