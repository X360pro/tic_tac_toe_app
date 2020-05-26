package components

import (
	"testing"
)

func TestCell(t *testing.T) {
	var list = []struct {
		name     string
		mark     string
		expected Player
	}{
		{"abhishek", "X", Player{myMark: "X", Name: "abhishek"}},
		{"ayushman", "O", Player{myMark: "O", Name: "ayushman"}},
		{"Tushhar", "X", Player{myMark: "X", Name: "Tushhar"}},
		{"Devesh", "X", Player{myMark: "X", Name: "Devesh"}},
		{"Eshan", "O", Player{myMark: "O", Name: "Eshan"}},
	}
	for _, str := range list {
		actual := CreatePlayer(str.name, str.mark)
		if *actual != str.expected {
			t.Error("In CreatePlayer Actual is :", *actual, "But expected is : ", str.expected)
		}
	}
}
