package main

func (s State) NextStates() []State {
	nextStates := make([]State, 0)
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if s.Board[r][c] == None {
				nextState := s.MakeMove(r, c)
				nextStates = append(nextStates, nextState)
			}
		}
	}
	return nextStates
}

func (s State) MakeMove(r int, c int) State {
	if s.Board[r][c] != None {
		panic("invalid move; piece already exist")
	}
	nextState := s
	nextState.WhoMovesNext = 3 - s.WhoMovesNext
	nextState.Board[r][c] = s.WhoMovesNext
	nextState.NumMovesSoFar++
	return nextState
}
