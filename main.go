package main

import (
	"fmt"
	"math/rand"

	gocheckers "github.com/fvolcic/gocheckers"
)

const (
	depth = 6
)

var searchedCount int

func findBestMove(board *gocheckers.CheckersBoard, blackPlayer bool) []int {

	moves := board.GenerateMoves()

	if len(moves) == 1 {
		return moves[0]
	}

	searchedCount = 0

	bestMove := moves[0]
	bestValue := minimax(board, depth, blackPlayer)

	fmt.Printf("Move: %v, Value: %d\n", bestMove, bestValue)

	for i := 1; i < len(moves); i++ {
		boardCopy := board.GenerateDeepCopy()
		boardCopy.MakeMove(moves[i])

		value := minimax(boardCopy, depth, blackPlayer)

		if blackPlayer {
			if value > bestValue {
				bestValue = value
				bestMove = moves[i]
			}
		} else {
			if value < bestValue {
				bestValue = value
				bestMove = moves[i]
			}
		}
	}
	fmt.Printf("Searched %d positions\n", searchedCount)
	return bestMove
}

func evaluate(board *gocheckers.CheckersBoard) int {
	raw := board.GetBoardData()

	eval := 0

	blackPieceCount := 0
	blackKingCount := 0

	whitePieceCount := 0
	whiteKingCount := 0

	for row := 0; row < len(raw); row++ {
		for col := 0; col < len(raw[row]); col++ {
			switch raw[row][col] {
			case gocheckers.Black:
				blackPieceCount++
			case gocheckers.White:
				whitePieceCount++
			case gocheckers.BlackKing:
				blackKingCount++
				blackPieceCount++
			case gocheckers.WhiteKing:
				whiteKingCount++
				whitePieceCount++
			}
		}
	}

	eval += 100*(blackPieceCount-whitePieceCount) + 200*(blackKingCount-whiteKingCount)

	return eval
}

func minimax(board *gocheckers.CheckersBoard, depth int, maximizingPlayer bool) int {

	searchedCount++

	if depth == 0 {
		return evaluate(board)
	}

	if maximizingPlayer {
		bestValue := -1000000
		moves := board.GenerateMoves()

		for i := 0; i < len(moves); i++ {
			board.MakeMove(moves[i])
			boardCopy := board.GenerateDeepCopy()
			value := minimax(boardCopy, depth-1, false)
			bestValue = max(bestValue, value)
		}

		return bestValue
	} else {
		bestValue := 1000000
		moves := board.GenerateMoves()

		for i := 0; i < len(moves); i++ {
			board.MakeMove(moves[i])
			boardCopy := board.GenerateDeepCopy()
			value := minimax(boardCopy, depth-1, true)
			bestValue = min(bestValue, value)
		}

		return bestValue
	}
}

func main() {

	board := gocheckers.NewCheckersBoard()
	for {

		fmt.Println(board.ToString())

		if board.GetTurn() == 1 {
			fmt.Println("Black's turn")

			// fmt.Println("black is thinking...")

			// move := findBestMove(board, true)
			// board.MakeMove(move)
			// continue
		} else {

			fmt.Println("White is thinking...")

			move := findBestMove(board, false)
			board.MakeMove(move)
			continue
		}

		moves := board.GenerateMoves()

		for i := 0; i < len(moves); i++ {
			fmt.Printf("%d.) ", i)
			fmt.Println(moves[i])
		}

		fmt.Println(board.ToString())

		whiteCount := board.GetPieceCount(gocheckers.White) + board.GetPieceCount(gocheckers.WhiteKing)
		blackCount := board.GetPieceCount(gocheckers.Black) + board.GetPieceCount(gocheckers.BlackKing)

		fmt.Println("White: ", whiteCount)
		fmt.Println("Black: ", blackCount)
		fmt.Println()

		fmt.Println("   32     31    30     29")
		fmt.Println("28     27    26    25")
		fmt.Println("   24     23    22     21")
		fmt.Println("20     19    18    17")
		fmt.Println("   16     15    14     13")
		fmt.Println("12     11    10     9")
		fmt.Println("    8      7     6      5")
		fmt.Println(" 4      3     2     1")

		fmt.Println("\nEnter your move: ")

		var move int

		//fmt.Scan(&move)

		move = rand.Intn(len(moves))

		success := board.MakeMove(moves[move])

		fmt.Println("Success: ", success)

	}

}
