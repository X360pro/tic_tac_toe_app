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
	b.Cells[position].SetMark(mark)
	for i := 0; i < 9; i++ {
		b.Cells[i].GetMark()
		if i%3 == 2 {
			fmt.Print("\n")
		}
	}
	fmt.Print("\n")
}
