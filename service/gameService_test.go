package service

import (
	"components"
	"errors"
	"testing"
)

func TestPlay(t *testing.T) {
	var list = []struct {
		position uint8
		g        *GameService
		expected error
	}{
		{9, &GameService{&ResultService{&BoardService{&components.Board{
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
			}}}}, &components.Player{Mark: "X", Name: "yo"}, &components.Player{Mark: "O", Name: "yo"}}, errors.New("Positoin is out of bounds")},
		{4, &GameService{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			}}}}, &components.Player{Mark: "X", Name: "yo"}, &components.Player{Mark: "O", Name: "yo"}}, errors.New("Positoin is out of bounds")},
	}

	for _, str := range list {
		err, _ := str.g.Play(str.position, &components.Player{Name: "nil", Mark: "X"})
		if err.Error() != str.expected.Error() {
			t.Error("fail")
		}
	}
}
