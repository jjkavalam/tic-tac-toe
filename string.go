package main

import (
	"bytes"
	"fmt"
)

func (s State) String() string {
	b := &bytes.Buffer{}

	boardSymbols := map[int]string{
		None:     "-",
		Me:       "x",
		Opponent: "o",
	}

	_, err := fmt.Fprintf(b, "next = %s, moves = %d", boardSymbols[s.WhoMovesNext], s.NumMovesSoFar)
	for r := 0; r < 3; r++ {
		_, err = fmt.Fprintln(b)
		for c := 0; c < 3; c++ {
			_, err = fmt.Fprintf(b, "%s ", boardSymbols[s.Board[r][c]])
		}
	}

	if err != nil {
		return err.Error()
	}

	return b.String()
}
