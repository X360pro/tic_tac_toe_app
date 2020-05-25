package components

import (
	"cell"
	"testing"
)

func TestCell(t *testing.T) {
	var list = []struct {
		mark     string
		expected cell.Cell
	}{
		{cell.XMark, cell.Cell{Mark: cell.XMark}},
		{cell.OMark, cell.Cell{Mark: cell.OMark}},
		{cell.NoMark, cell.Cell{Mark: cell.NoMark}},
	}
	for _, str := range list {
		actual := cell.NewCell()
		if actual.Mark != cell.NoMark {
			t.Error("In NewCell Actual is :", *actual, "But expected is : ", cell.NoMark)
		}
		_ = cell.SetMark(actual, str.mark)
		if *actual != str.expected {
			t.Error("In SetMark Actual is :", *actual, "But expected is : ", str.expected)
		}
	}
}
