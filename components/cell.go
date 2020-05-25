package components

import "fmt"

type Cell struct {
	Mark string
}

const (
	NoMark = "-"
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{Mark: NoMark}
}

func (c *Cell) GetMark() {
	fmt.Print(" ", c.Mark, " ")
}
func (c *Cell) SetMark(mark string) bool {
	if c.Mark != NoMark {
		fmt.Println("Cell is not empty")
		return false
	} else {
		c.Mark = mark
		return true
	}
}
