package main

import "fmt"

const Winning = 1
const Losing = 2
const Drawing = 3

func (s State) GetComputerMove() State {
	if s.WhoMovesNext != Me {
		panic("it is not my turn")
	}
	t, s := s.Type()
	types := map[int]string{
		Winning: "Winning",
		Losing:  "Losing",
		Drawing: "Drawing",
	}
	fmt.Println("Found move of type", types[t])
	fmt.Println(s)
	if t == Winning {
		fmt.Println("I will win no matter what !")
	} else if t == Losing {
		fmt.Println("I accept defeat")
	}
	return s
}

func (s State) Type() (int, State) {
	// base case
	// if already won / lost / drawn
	// then the next state is the same
	if s.HasWon() {
		return Winning, s
	} else if s.HasLost() {
		return Losing, s
	} else if s.IsFull() {
		return Drawing, s
	}

	if s.WhoMovesNext != Me {
		panic("ask me only when it is my move next")
	}

	// let us look at each possible move
	nextStates := s.NextStates()

	var losingMove State

	var foundDraw bool
	var drawingMove State

	var foundWin bool
	var winningMove State

	if s.NumMovesSoFar == 8 {
		for _, s := range nextStates {
			// there are no opponent moves to check
			outcome, _ := s.Type()

			switch outcome {
			case Winning:
				foundWin = true
				winningMove = s
			case Drawing:
				foundDraw = true
				drawingMove = s
			case Losing:
				losingMove = s
			}
		}
	} else {
		for _, s := range nextStates {

			// for each move, let us play the opponent
			opponentMoves := s.NextStates()

			opponentCouldDraw := false
			opponentCanWin := false

			// i.e. from all possible moves the opponent can make
			// check if there moves where we will lose
			// if there moves where he can draw, record that also
			for _, s2 := range opponentMoves {
				outcome, _ := s2.Type()
				switch outcome {
				case Losing:
					opponentCanWin = true
				case Drawing:
					opponentCouldDraw = true
				}
			}

			// the opponent will surely try to win, if not at least try to draw
			if opponentCanWin {
				// since there is a way for the opponent to guarantee his victory
				// we should not make this move (unless we are accepting defeat)
				losingMove = s
				continue
			} else if opponentCouldDraw {
				// there is no way the opponent can win, but he can still draw it
				// So it is not ideal, but we will still keep it as an option
				foundDraw = true
				drawingMove = s
			} else {
				// there is no way the opponent can win or draw this;
				// now, that's the move we want to make !
				foundWin = true
				winningMove = s
			}
		}
	}

	// if we found a winning move, go for it;
	// it not, go for the drawing move
	// if neither, what is else is left to do, accept defeat !
	if foundWin {
		return Winning, winningMove
	} else if foundDraw {
		return Drawing, drawingMove
	} else {
		return Losing, losingMove
	}

}
