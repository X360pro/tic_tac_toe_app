package service

import (
	"testing"
	"tictactoe/components"
)

func TestCheckRow(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		pos      uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, 0, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, 6, true},
		{&ResultService{&BoardService{&components.Board{
			Size: 4,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},

				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},

				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.XMark},

				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, 10, true},
	}
	for _, test := range tests {
		if test.input.checkRow(test.mark, test.pos) != test.expected {
			t.Error("check row failed")
		}
	}
}
func TestCheckColumn(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		pos      uint8
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, 0, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, 5, true},
		{&ResultService{&BoardService{&components.Board{
			Size: 4,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},

				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},

				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},

				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.OMark, 9, true},
	}
	for _, test := range tests {
		if test.input.checkColumn(test.mark, test.pos) != test.expected {
			t.Error("check column failed")
		}
	}
}

func TestCheckLRDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.OMark, true},
		{&ResultService{&BoardService{&components.Board{
			Size: 4,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},

				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},

				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},

				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.checkLRDiagonal(test.mark) != test.expected {
			t.Error("check LR Diagonal failed", test.input.Size)
		}
	}
}
func TestCheckRLDiagonal(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.OMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, true},
		{&ResultService{&BoardService{&components.Board{
			Size: 4,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},

				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},

				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.OMark},

				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.checkRLDiagonal(test.mark) != test.expected {
			t.Error("check RL Diagonal failed", test.input.Size)
		}
	}
}
