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

func (goban *Goban) UnCapture(positions [][]int32, color int8) {
	if len(positions) == 0 {
		return
	}
	for _, pos := range positions {
		goban.SetElem(pos[0], pos[1], color)
		Gomoku.ActivePlayer.AddCaptured(-1)
	}
}

func (goban *Goban) tryCaptureUp(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topElem := goban.GetTopElem(row, col)
		if topElem == 0 {
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

func (goban *Goban) tryCaptureDown(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomElem := goban.GetBottomElem(row, col)
		if bottomElem == 0 {
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

func (goban *Goban) tryCaptureLeft(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		leftElem := goban.GetLeftElem(row, col)
		if leftElem == 0 {
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

func (goban *Goban) tryCaptureRight(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		rightElem := goban.GetRightElem(row, col)
		if rightElem == 0 {
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

func (goban *Goban) tryCaptureTopLeft(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topLeftElem := goban.GetTopLeftElem(row, col)
		if topLeftElem == 0 {
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

func (goban *Goban) tryCaptureTopRight(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		topRightElem := goban.GetTopRightElem(row, col)
		if topRightElem == 0 {
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

func (goban *Goban) tryCaptureBottomLeft(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomLeftElem := goban.GetBottomLeftElem(row, col)
		if bottomLeftElem == 0 {
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

func (goban *Goban) tryCaptureBottomRight(row int32, col int32, currentColor int8, opponentColor int8) [][]int32 {
	for idx := 0; idx < 3; idx++ {
		bottomRightElem := goban.GetBottomRightElem(row, col)
		if bottomRightElem == 0 {
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

func (goban *Goban) CanBeCaptured(row int32, col int32, currentColor int8) bool {
	if goban.canBeCapturedVertical(row, col, currentColor) ||
		goban.canBeCapturedHorizontal(row, col, currentColor) ||
		goban.canBeCapturedDiagonal_1(row, col, currentColor) ||
		goban.canBeCapturedDiagonal_2(row, col, currentColor) {
		return true
	}
	return false
}

func (goban *Goban) CountPossibleCaptures(row int32, col int32, currentColor int8) int32 {
	count := int32(0)
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

func (goban *Goban) canBeCapturedVertical(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for goban.GetTopElem(row, col) == currentColor {
		row--
	}
	upperCell := goban.GetTopElem(row, col)
	for goban.GetBottomElem(row, col) == currentColor {
		count++
		row++
	}
	bottomCell := goban.GetBottomElem(row, col)
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedHorizontal(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for goban.GetLeftElem(row, col) == currentColor {
		col--
	}
	upperCell := goban.GetLeftElem(row, col)
	for goban.GetRightElem(row, col) == currentColor {
		count++
		col++
	}
	bottomCell := goban.GetRightElem(row, col)
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedDiagonal_1(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for goban.GetTopLeftElem(row, col) == currentColor {
		col--
		row--
	}
	upperCell := goban.GetTopLeftElem(row, col)
	for goban.GetBottomRightElem(row, col) == currentColor {
		count++
		col++
		row++
	}
	bottomCell := goban.GetBottomRightElem(row, col)
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}

func (goban *Goban) canBeCapturedDiagonal_2(row int32, col int32, currentColor int8) bool {
	opponentColor := GetOpponentColor(currentColor)
	count := 1
	for goban.GetTopRightElem(row, col) == currentColor {
		col++
		row--
	}
	upperCell := goban.GetTopRightElem(row, col)
	for goban.GetBottomLeftElem(row, col) == currentColor {
		count++
		col--
		row++
	}
	bottomCell := goban.GetBottomLeftElem(row, col)
	if count == 2 && (upperCell == opponentColor && bottomCell == 0 || bottomCell == opponentColor && upperCell == 0) {
		return true
	}
	return false
}
