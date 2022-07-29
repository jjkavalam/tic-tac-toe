package main

func main() {
	s := EmptyState()
	s.GetComputerMove().MakeOpponentMove(0, 0).GetComputerMove().MakeOpponentMove(2, 1).GetComputerMove().MakeOpponentMove(1, 1).GetComputerMove()
}
