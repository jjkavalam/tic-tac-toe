package main

func (s State) MakeOpponentMove(r int, c int) State {
	if s.WhoMovesNext != Opponent {
		panic("it is not opponent's turn")
	}
	return s.MakeMove(r, c)
}
