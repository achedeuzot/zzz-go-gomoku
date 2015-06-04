package arena

func (goban *Goban) Capture(row int32, col int32) [][]int32 {
	currentColor := Gomoku.ActivePlayer.GetColor()
	opponentColor := Gomoku.OtherPlayer.GetColor()
	capturedPawns := make([][]int32, 0)
	capturedPawns = append(capturedPawns, goban.tryCaptureUp(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureDown(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureLeft(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureRight(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureTopLeft(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureTopRight(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureBottomLeft(row, col, currentColor, opponentColor)...)
	capturedPawns = append(capturedPawns, goban.tryCaptureBottomRight(row, col, currentColor, opponentColor)...)
	Gomoku.ActivePlayer.AddCaptured(int8(len(capturedPawns)))
	return capturedPawns
}

func (goban *Goban) UnCapture(positions [][]int32, color byte) {
	if len(positions) == 0 {
		return
	}
	for _, pos := range positions {
		goban.SetElem(pos[0], pos[1], color)
		Gomoku.ActivePlayer.AddCaptured(-1)
	}
}

func (goban *Goban) tryCaptureUp(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topElem, ok := goban.GetTopElem(row, col)
		if ok == false || topElem == 0 {
			return nil
		} else if idx < 2 && topElem != opponentColor {
			return nil
		} else if idx == 2 && topElem != currentColor {
			return nil
		}
		row--
	}
	goban.SetElem(row+1, col, 0)
	goban.SetElem(row+2, col, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row + 1
	pos[0][1] = col
	pos[1][0] = row + 2
	pos[1][1] = col
	return pos
}

func (goban *Goban) tryCaptureDown(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomElem, ok := goban.GetBottomElem(row, col)
		if ok == false || bottomElem == 0 {
			return nil
		} else if idx < 2 && bottomElem != opponentColor {
			return nil
		} else if idx == 2 && bottomElem != currentColor {
			return nil
		}
		row++
	}
	goban.SetElem(row-1, col, 0)
	goban.SetElem(row-2, col, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row - 1
	pos[0][1] = col
	pos[1][0] = row - 2
	pos[1][1] = col
	return pos
}

func (goban *Goban) tryCaptureLeft(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		leftElem, ok := goban.GetLeftElem(row, col)
		if ok == false || leftElem == 0 {
			return nil
		} else if idx < 2 && leftElem != opponentColor {
			return nil
		} else if idx == 2 && leftElem != currentColor {
			return nil
		}
		col--
	}
	goban.SetElem(row, col+1, 0)
	goban.SetElem(row, col+2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row
	pos[0][1] = col + 1
	pos[1][0] = row
	pos[1][1] = col + 2
	return pos
}

func (goban *Goban) tryCaptureRight(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		rightElem, ok := goban.GetRightElem(row, col)
		if ok == false || rightElem == 0 {
			return nil
		} else if idx < 2 && rightElem != opponentColor {
			return nil
		} else if idx == 2 && rightElem != currentColor {
			return nil
		}
		col++
	}
	goban.SetElem(row, col-1, 0)
	goban.SetElem(row, col-2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row
	pos[0][1] = col - 1
	pos[1][0] = row
	pos[1][1] = col - 2
	return pos
}

func (goban *Goban) tryCaptureTopLeft(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topLeftElem, ok := goban.GetTopLeftElem(row, col)
		if ok == false || topLeftElem == 0 {
			return nil
		} else if idx < 2 && topLeftElem != opponentColor {
			return nil
		} else if idx == 2 && topLeftElem != currentColor {
			return nil
		}
		row--
		col--
	}
	goban.SetElem(row+1, col+1, 0)
	goban.SetElem(row+2, col+2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row + 1
	pos[0][1] = col + 1
	pos[1][0] = row + 2
	pos[1][1] = col + 2
	return pos
}

func (goban *Goban) tryCaptureTopRight(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topRightElem, ok := goban.GetTopRightElem(row, col)
		if ok == false || topRightElem == 0 {
			return nil
		} else if idx < 2 && topRightElem != opponentColor {
			return nil
		} else if idx == 2 && topRightElem != currentColor {
			return nil
		}
		row--
		col++
	}
	goban.SetElem(row+1, col-1, 0)
	goban.SetElem(row+2, col-2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row + 1
	pos[0][1] = col - 1
	pos[1][0] = row + 2
	pos[1][1] = col - 2
	return pos
}

func (goban *Goban) tryCaptureBottomLeft(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomLeftElem, ok := goban.GetBottomLeftElem(row, col)
		if ok == false || bottomLeftElem == 0 {
			return nil
		} else if idx < 2 && bottomLeftElem != opponentColor {
			return nil
		} else if idx == 2 && bottomLeftElem != currentColor {
			return nil
		}
		row++
		col--
	}
	goban.SetElem(row-1, col+1, 0)
	goban.SetElem(row-2, col+2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row - 1
	pos[0][1] = col + 1
	pos[1][0] = row - 2
	pos[1][1] = col + 2
	return pos
}

func (goban *Goban) tryCaptureBottomRight(row int32, col int32, currentColor byte, opponentColor byte) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomRightElem, ok := goban.GetBottomRightElem(row, col)
		if ok == false || bottomRightElem == 0 {
			return nil
		} else if idx < 2 && bottomRightElem != opponentColor {
			return nil
		} else if idx == 2 && bottomRightElem != currentColor {
			return nil
		}
		row++
		col++
	}
	goban.SetElem(row-1, col-1, 0)
	goban.SetElem(row-2, col-2, 0)
	pos := make([][]int32, 2)
	pos[0] = make([]int32, 2)
	pos[1] = make([]int32, 2)
	pos[0][0] = row - 1
	pos[0][1] = col - 1
	pos[1][0] = row - 2
	pos[1][1] = col - 2
	return pos
}

func (goban *Goban) CanBeCaptured(row int32, col int32, currentColor byte) bool {
	if goban.canBeCapturedVertical(row, col, currentColor) ||
		goban.canBeCapturedHorizontal(row, col, currentColor) ||
		goban.canBeCapturedDiagonal_1(row, col, currentColor) ||
		goban.canBeCapturedDiagonal_2(row, col, currentColor) {
		return true
	}
	return false
}

func (goban *Goban) CountPossibleCaptures(row int32, col int32, currentColor byte) (count int32) {
	if goban.canBeCapturedVertical(row, col, currentColor) {
		count++
	}
	if goban.canBeCapturedHorizontal(row, col, currentColor) {
		count++
	}
	if goban.canBeCapturedDiagonal_1(row, col, currentColor) {
		count++
	}
	if goban.canBeCapturedDiagonal_2(row, col, currentColor) {
		count++
	}
	return count
}

func (goban *Goban) canBeCapturedVertical(row int32, col int32, currentColor byte) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for val, _ := goban.GetTopElem(row, col); val == currentColor; val, _ = goban.GetTopElem(row, col) {
		row--
	}
	upperCell, okUp := goban.GetTopElem(row, col)
	for val, _ := goban.GetBottomElem(row, col); val == currentColor; val, _ = goban.GetBottomElem(row, col) {
		count++
		row++
	}
	bottomCell, okBottom := goban.GetBottomElem(row, col)
	if okUp == false || okBottom == false {
		return false
	}
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedHorizontal(row int32, col int32, currentColor byte) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for val, _ := goban.GetLeftElem(row, col); val == currentColor; val, _ = goban.GetLeftElem(row, col) {
		col--
	}
	upperCell, okUp := goban.GetLeftElem(row, col)
	for val, _ := goban.GetRightElem(row, col); val == currentColor; val, _ = goban.GetRightElem(row, col) {
		count++
		col++
	}
	bottomCell, okBottom := goban.GetRightElem(row, col)
	if okUp == false || okBottom == false {
		return false
	}
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedDiagonal_1(row int32, col int32, currentColor byte) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for val, _ := goban.GetTopLeftElem(row, col); val == currentColor; val, _ = goban.GetTopLeftElem(row, col) {
		col--
		row--
	}
	upperCell, okUp := goban.GetTopLeftElem(row, col)
	for val, _ := goban.GetBottomRightElem(row, col); val == currentColor; val, _ = goban.GetBottomRightElem(row, col) {
		count++
		col++
		row++
	}
	bottomCell, okBottom := goban.GetBottomRightElem(row, col)
	if okUp == false || okBottom == false {
		return false
	}
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedDiagonal_2(row int32, col int32, currentColor byte) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for val, _ := goban.GetTopRightElem(row, col); val == currentColor; val, _ = goban.GetTopRightElem(row, col) {
		col++
		row--
	}
	upperCell, okUp := goban.GetTopRightElem(row, col)
	for val, _ := goban.GetBottomLeftElem(row, col); val == currentColor; val, _ = goban.GetBottomLeftElem(row, col) {
		count++
		col--
		row++
	}
	bottomCell, okBottom := goban.GetBottomLeftElem(row, col)
	if okUp == false || okBottom == false {
		return false
	}
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}
