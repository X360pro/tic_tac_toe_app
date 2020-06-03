package service

import (
	"log"
	"testing"
	"tictactoe/components"
)

func TestPlay(t *testing.T) {
	var list = []struct {
		position uint8
		g        *GameService
		err      string
	}{
		{3, &GameService{&ResultService{NewBoardService(3)}, [2]*components.Player{
			{Name: "abhishek", Mark: components.OMark},
			{Name: "soham", Mark: components.XMark},
		}}, "nil"},
		{7, &GameService{&ResultService{NewBoardService(3)}, [2]*components.Player{
			{Name: "abhishek", Mark: components.OMark},
			{Name: "soham", Mark: components.XMark},
		}}, "nil"},
		{4, &GameService{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			}}}}, [2]*components.Player{
			{Name: "abhishek", Mark: components.OMark},
			{Name: "soham", Mark: components.XMark},
		}}, "Positoin is out of bounds"},
	}

	for _, str := range list {
		err, _ := str.g.Play(str.position)
		if err != nil {
			if err.Error() != str.err {
				t.Error("errors didn't matched", str.position)
			}
		} else if str.g.Cells[str.position].GetMark() != str.g.player[0].Mark {
			t.Error("mark not matching")
		}
		err, _ = str.g.Play(str.position + 1)
		if err != nil {
			if err.Error() != str.err {
				log.Fatal("errors didn't matched")
			}
		} else if str.g.Cells[str.position+1].GetMark() != str.g.player[1].Mark {
			t.Error("mark not matching")
		}
	}
}
