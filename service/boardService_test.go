package service

import (
	"components"
	"testing"
)

func TestPutMarkInPosition(t *testing.T) {
	var list = []struct {
		position uint8
		expected BoardService
		pl       components.Player
	}{
		{3, BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		}}, components.Player{Name: "admam", Mark: "X"}},
		{6, BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		}}, components.Player{Name: "admam", Mark: "X"}},
	}
	for _, str := range list {
		str.expected.PutMarkInPosition(str.position, &str.pl)
		if str.expected.Cells[str.position].GetMark() != str.pl.Mark {
			t.Error("Put mark in position failed")
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	var list = []struct {
		b        BoardService
		expected bool
	}{
		{BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			}}}, false},
		{BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
			}}}, true},
	}
	for _, str := range list {

		if str.b.CheckBoardIsFull() != str.expected {
			t.Error("Check board is full failed")
		}
	}
}
