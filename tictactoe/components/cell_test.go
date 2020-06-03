package components

import (
	"testing"
)

func TestCell(t *testing.T) {
	var list = []struct {
		mark     string
		expected Cell
	}{
		{XMark, Cell{Mark: XMark}},
		{OMark, Cell{Mark: OMark}},
		{NoMark, Cell{Mark: NoMark}},
	}
	for _, str := range list {
		actual := NewCell()
		if actual.Mark != NoMark {
			t.Error("In NewCell Actual is :", *actual, "But expected is : ", NoMark)
		}
		_ = actual.SetMark(str.mark)
		if *actual != str.expected {
			t.Error("In SetMark Actual is :", *actual, "But expected is : ", str.expected)
		}
	}
	for _, str := range list {
		if str.mark != str.expected.GetMark() {
			t.Error("In NewCell Actual is :", str.expected.GetMark(), "But expected is : ", str.mark)
		}
	}
}
