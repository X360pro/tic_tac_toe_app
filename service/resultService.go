package service

type ResultService struct {
	*BoardService
}

func NewResultService(boardService *BoardService) *ResultService {
	return &ResultService{boardService}
}

func (r *ResultService) CheckRow(mark string) bool {
	//for rows
	count := 0
	for i := 0; i < int(r.Size*r.Size); i++ {
		if count == int(r.Size) {
			return true
		}
		if r.Cells[i].GetMark() == mark {
			count++
		}
		if i+1%int(r.Size) == 0 {
			count = 0
		}
	}
	return false
}
func (r *ResultService) CheckColumn(mark string) bool {
	//for columns
	count_col := make([]int, r.Size)
	for i := 0; i < int(r.Size*r.Size); i++ {

		if r.Cells[i].GetMark() == mark {
			count_col[i%int(r.Size)]++
		}
		if contains(count_col, int(r.Size)) {
			return true
		}
	}
	return false
}
func (r *ResultService) CheckLRDiagonal(mark string) bool {
	//for LR diagonal
	count := 0
	j := 0
	for i := 0; i < int(r.Size); i++ {

		if r.Cells[(i*int(r.Size))+j].GetMark() == mark {
			count++
		}
		if count == int(r.Size) {
			return true
		}
		j++
	}
	return false
}
func (r *ResultService) CheckRLDiagonal(mark string) bool {
	//for RL diagonal
	count := 0
	for i := 0; i < int(r.Size); i++ {

		if r.Cells[(i+1)*(int(r.Size)-1)].GetMark() == mark {
			count++
		}
		if count == int(r.Size) {
			return true
		}
	}
	return false
}
func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
