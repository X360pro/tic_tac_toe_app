package service

import (
	"tictactoe/components"
)

type ResultService struct {
	*BoardService
}

func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (r *ResultService) checkRow(mark string, pos uint8) bool {
	start := r.Size * (pos / r.Size)
	for i := start; i < start+r.Size; i++ {
		if r.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}
func (r *ResultService) checkColumn(mark string, pos uint8) bool {
	start := pos % r.Size
	for i := start; i < r.Size*r.Size; i += r.Size {
		if r.Cells[i].GetMark() != mark {
			return false
		}
	}
	return true
}
func (r *ResultService) checkLRDiagonal(mark string) bool {
	j := 0
	for i := 0; i < int(r.Size*r.Size); i = ((j * int(r.Size)) + j) {
		if r.Cells[i].GetMark() != mark {
			return false
		}
		j++
	}
	return true
}
func (r *ResultService) checkRLDiagonal(mark string) bool {
	j := 1
	for i := r.Size - 1; i <= r.Size*(r.Size-1); i = (uint8(j) * (r.Size - 1)) {
		if r.Cells[i].GetMark() != mark {
			return false
		}
		j++
	}
	return true
}
func (r *ResultService) GiveResult(player *components.Player, pos uint8) string {

	if r.checkRow(player.Mark, pos) {
		return "win"
	} else if r.checkColumn(player.Mark, pos) {
		return "win"
	} else if r.checkLRDiagonal(player.Mark) {
		return "win"
	} else if r.checkRLDiagonal(player.Mark) {
		return "win"
	} else if r.CheckBoardIsFull() {
		return "draw"
	}
	return "ongoing"
}
