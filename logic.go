package main

import "fmt"

const UNKNOWN = 0
const Winning = 1
const Losing = 2
const Drawing = 3

var predictions = map[int]string{
	Winning: "Winning",
	Losing:  "Losing",
	Drawing: "Drawing",
}

func (s State) GetComputerMove() State {
	if s.WhoMovesNext != Me {
		panic("it is not my turn")
	}
	p, s := s.Analyze()
	fmt.Println("Found move with predicted:", predictions[p])
	fmt.Println(s)
	if p == Winning {
		fmt.Println("I will win no matter what !")
	} else if p == Losing {
		fmt.Println("I accept defeat")
	}
	return s
}

// Analyze thinks like a player who wants to win and evaluates whether the current state can lead to a
// Winning / Losing / Drawing game for Me, assuming the Opponent plays his best game.
// Since the analysis invariably involves considering the various moves, it can also return the new State
// after the best possible move has been made.
//
// Analyze is expected to be called when it is player Me's turn.
//
// If the board is already won/lost/full it returns immediately and the new state is simply the current state itself.
// In this case, the returned prediction is also the actual outcome.
//
// To compute the "best" move, we look at all possible moves Me can make in the given state.
// For each of these moves, we try out each possible move that the Opponent can make.
// We Analyze (recursively) the resulting State to evaluate how good our original move was.
// The rules of evaluation are as follows:
// - if it is somehow possible for the  Opponent to put the game into a Losing situation; then the move is bad.
// - if the Opponent can never put the game into a Losing situation; but can make it Drawing; then the move could be considered.
// - if the Opponent can never make the game Losing or Drawing; then we have a winning move, and we will take it !
//
// The above is essentially the Minimax algorithm.
//
func (s State) Analyze() (int, State) {
	if s.WhoMovesNext != Me {
		panic("ask me only when it is my move next")
	}

	// base case
	// if already won / lost / drawn
	// then the next state is the same
	outcome := s.GetOutcome()
	if outcome != UNKNOWN {
		return outcome, s
	}

	// let us look at each possible move
	nextStates := s.NextStates()

	if s.NumMovesSoFar == 8 {
		// if 8 moves have been made already, that means there is only one more move left in this game
		// Hence, there is actually one possible move that we can make;
		// and the resulting position can be directly analyzed.
		nextState := nextStates[0]
		return nextState.GetOutcome(), nextState
	}

	var losingMove State

	var foundDraw bool
	var drawingMove State

	var foundWin bool
	var winningMove State

	for _, s := range nextStates {
		// for each move, let us play the opponent
		opponentMoves := s.NextStates()

		opponentCanDraw := false
		opponentCanWin := false

		// i.e. from all possible moves the opponent can make
		// check if there are any moves that guarantee our defeat.
		// Also, if there are moves where he can draw, take note of that also.
		for _, s2 := range opponentMoves {
			outcome, _ := s2.Analyze()
			switch outcome {
			case Losing:
				opponentCanWin = true
			case Drawing:
				opponentCanDraw = true
			}
		}

		// the opponent will surely try to win, if not at least try to draw
		if opponentCanWin {
			// since there is a way for the opponent to guarantee his victory
			// s is a bad move (unless we have no choice, but to accept defeat)
			losingMove = s
			continue
		} else if opponentCanDraw {
			// there is no way the opponent can win; but he can still draw it.
			// So it is not ideal, but we will still keep it as an option
			foundDraw = true
			drawingMove = s
		} else {
			// there is no way the opponent can win or draw this;
			// now, this is the kind of move we want to make !
			foundWin = true
			winningMove = s
		}
	}

	// if we found a winning move, go for it;
	// it not, go for the drawing move
	// if neither, what else is left to do, accept defeat !
	if foundWin {
		return Winning, winningMove
	} else if foundDraw {
		return Drawing, drawingMove
	} else {
		return Losing, losingMove
	}

}
