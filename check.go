package main

func (s State) HasWon() bool {
	return s.HasRowColOrDiagonal(Me)
}

func (s State) HasLost() bool {
	return s.HasRowColOrDiagonal(Opponent)
}

func (s State) HasRowColOrDiagonal(piece int) bool {
	// for each row, sweep the columns
	for r := 0; r < 3; r++ {
		count := 0
		for c := 0; c < 3; c++ {
			if s.Board[r][c] == piece {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	// for each column, sweep the rows
	for c := 0; c < 3; c++ {
		count := 0
		for r := 0; r < 3; r++ {
			if s.Board[r][c] == piece {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	// check the diagonal
	count := 0
	for d := 0; d < 3; d++ {
		if s.Board[d][d] == piece {
			count++
		}
	}
	if count == 3 {
		return true
	}
	// check the off diagonal
	count = 0
	for d := 0; d < 3; d++ {
		if s.Board[d][2-d] == piece {
			count++
		}
	}
	if count == 3 {
		return true
	}
	return false
}
