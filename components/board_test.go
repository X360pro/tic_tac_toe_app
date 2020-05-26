package components

import (
	"testing"
)

func TestBoard(t *testing.T) {
	var list = []struct {
		board_size           uint8
		expexted_no_of_cells int
	}{
		{3, 9},
		{4, 16},
		{5, 25},
	}
	for _, str := range list {
		actual := NewBoard(str.board_size)
		if len(actual.Cells) != str.expexted_no_of_cells {
			t.Error("In NewBoard Actual no. of cells is :", len(actual.Cells), "But expected is : ", str.expexted_no_of_cells)
		}
	}
}
