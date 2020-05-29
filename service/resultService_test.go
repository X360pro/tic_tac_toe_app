package service

import (
	"components"
	"testing"
)

func TestCheckRow(t *testing.T) {
	tests := []struct {
		input    *ResultService
		mark     string
		expected bool
	}{
		{&ResultService{&BoardService{&components.Board{
			Size: 2,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, true},

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
		}, components.OMark, true},
	}
	for _, test := range tests {
		if test.input.CheckRow(test.mark) != test.expected {
			t.Error("check row failed")
		}
	}
}

func TestCheckColumn(t *testing.T) {
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
				{Mark: components.XMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
			},
		},
		},
		}, components.OMark, true},
	}
	for _, test := range tests {
		if test.input.CheckColumn(test.mark) != test.expected {
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
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.CheckLRDiagonal(test.mark) != test.expected {
			t.Error("check LR diagonal failed")
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
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
			},
		},
		},
		}, components.OMark, true},

		{&ResultService{&BoardService{&components.Board{
			Size: 3,
			Cells: []*components.Cell{
				{Mark: components.NoMark},
				{Mark: components.OMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.NoMark},
				{Mark: components.XMark},
				{Mark: components.OMark},
				{Mark: components.OMark},
			},
		},
		},
		}, components.XMark, true},
	}
	for _, test := range tests {
		if test.input.CheckRLDiagonal(test.mark) != test.expected {
			t.Error("check RL diagonal failed")
		}
	}
}