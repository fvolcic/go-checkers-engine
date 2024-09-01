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

	moveValues := make([]int, len(moves))
	var moveValueChannels []chan int

	for i := 0; i < len(moves); i++ {
		moveValueChannels = append(moveValueChannels, make(chan int))
	}

	for i := 0; i < len(moves); i++ {
		boardCopy := board.GenerateDeepCopy()
		boardCopy.MakeMove(moves[i])
		go minimax(boardCopy, depth, blackPlayer, moveValueChannels[i])
	}

	for i := 0; i < len(moves); i++ {
		moveValues[i] = <-moveValueChannels[i]
	}

	bestMoveIndex := 0
	bestMoveValue := -1000000

	for i := 0; i < len(moves); i++ {
		if moveValues[i] > bestMoveValue {
			bestMoveValue = moveValues[i]
			bestMoveIndex = i
		}
	}

	fmt.Println("Searched count: ", searchedCount)
	searchedCount = 0

	return moves[bestMoveIndex]

}

func evaluate(board *gocheckers.CheckersBoard) int {
	raw := board.GetUnpaddedBoard()

	eval := 0

	blackPieceCount := 0
	blackKingCount := 0

	whitePieceCount := 0
	whiteKingCount := 0

	for row := 0; row < len(raw); row++ {
		for col := 0; col < len(raw[row]); col++ {
			switch raw[row][col] {
			case 1:
				blackPieceCount++
			case 2:
				whitePieceCount++
			case 3:
				blackKingCount++
				blackPieceCount++
			case 4:
				whiteKingCount++
				whitePieceCount++
			}
		}
	}

	eval += 100*(blackPieceCount-whitePieceCount) + 200*(blackKingCount-whiteKingCount)

	return eval
}

func minimax(board *gocheckers.CheckersBoard, depth int, maximizingPlayer bool, valueCh chan int) int {

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
			value := minimax(boardCopy, depth-1, false, nil)
			bestValue = max(bestValue, value)
		}

		if valueCh != nil {
			valueCh <- bestValue
		}

		return bestValue
	} else {
		bestValue := 1000000
		moves := board.GenerateMoves()

		for i := 0; i < len(moves); i++ {
			board.MakeMove(moves[i])
			boardCopy := board.GenerateDeepCopy()
			value := minimax(boardCopy, depth-1, true, nil)
			bestValue = min(bestValue, value)
		}

		if valueCh != nil {
			valueCh <- bestValue
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
