package main

const None = 0
const Me = 1
const Opponent = 2

type State struct {
	// WhoMovesNext is either 1 or 2
	WhoMovesNext  int
	Board         [3][3]int
	NumMovesSoFar int
}

func EmptyState() State {
	return State{
		WhoMovesNext: Me,
	}
}

func (s State) IsFull() bool {
	return s.NumMovesSoFar == 9
}
