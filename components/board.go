package components

type Board struct {
	Cells []*Cell
	Size  uint8
}

func NewBoard(size uint8) *Board {
	new := Board{}
	new.Cells = make([]*Cell, size*size)
	for i := range new.Cells {
		new.Cells[i] = NewCell()
	}
	new.Size = size
	return &new
}
