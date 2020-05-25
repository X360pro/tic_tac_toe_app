package components

import (
	"errors"
	"fmt"
)

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
func (c *Cell) SetMark(mark string) error {
	if c.Mark != NoMark {
		return errors.New("Cell is not empty")
	} else {
		c.Mark = mark
	}
	return nil
}
