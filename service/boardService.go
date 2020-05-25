package service

import (
	"components"
	"fmt"
)

type BoardService struct {
	*components.Board
}

func NewBoardService() *BoardService {
	return &BoardService{components.NewBoard(3)}
}

func (b *BoardService) Mark(position int, mark string) {
	err := b.Cells[position].SetMark(mark)
	if err != nil {
		fmt.Print(err)
	}
}
