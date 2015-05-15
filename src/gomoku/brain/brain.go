package brain

type Goban [361]uint8

func (goban *Goban) GetElem(row int, col int) uint8 {
	return goban[row*19+col]
}

func (goban *Goban) GetTopElem(row int, col int) uint8 {
	if row > 0 {
		return goban[(row-1)*19+col]
	}
	return -1
}

func (goban *Goban) GetBottomElem(row int, col int) uint8 {
	if row < 17 {
		return goban[(row+1)*19+col]
	}
	return -1
}

func (goban *Goban) GetLeftElem(row int, col int) uint8 {
	if col > 0 {
		return goban[row*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetRightElem(row int, col int) uint8 {
	if col < 17 {
		return goban[row*19+(col+1)]
	}
	return -1
}

func (goban *Goban) GetTopLeftElem(row int, col int) uint8 {
	if row > 0 && col > 0 {
		return goban[(row-1)*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetTopRightElem(row int, col int) uint8 {
	if row > 0 && col < 17 {
		return goban[(row-1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) GetBottomLeftElem(row int, col int) uint8 {
	if row < 17 && col > 0 {
		return goban[(row+1)*19+(col-1)]
	}
	return -1
}

func (goban *Goban) GetBottomRightElem(row int, col int) uint8 {
	if row < 17 && col < 17 {
		return goban[(row+1)*19+(col+1)]
	}
	return -1
}

func (goban *Goban) SetElem(row int, col int, val uint8) {
	goban[row*19+col] = val
}
