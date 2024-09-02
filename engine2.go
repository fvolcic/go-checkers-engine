package main

import (
	"fmt"
	"math"
	"math/rand"

	gocheckers "github.com/fvolcic/gocheckers"
)

const (
	Depth = 10
)

var alpha int = math.MaxInt // The best score that the minimizing player can guarantee
var beta int = math.MinInt  // The best score that the maximizing player can guarantee

var (
	searchCount = 0
)

func ResetAlphaBeta() {
	alpha = math.MaxInt
	beta = math.MinInt
}

func GetNextMove(board *gocheckers.CheckersBoard) []int {
	// Given our evaluation function,
	// black is the maximizing player and white is the minimizing player.

	searchCount = 0

	isMaximizing := true
	if board.GetTurn() == gocheckers.White {
		isMaximizing = false
	}

	moves := board.GenerateMoves()

	var bestValue int
	if isMaximizing {
		bestValue = math.MinInt
	} else {
		bestValue = math.MaxInt
	}

	var bestMove []int

	for _, move := range moves {
		boardCopy := board.GenerateDeepCopy()
		boardCopy.MakeMove(move)

		value := AlphaBeta(boardCopy, Depth, isMaximizing)

		if isMaximizing && value > bestValue {
			bestMove = move
			bestValue = value
		} else if !isMaximizing && value < bestValue {
			bestMove = move
			bestValue = value
		}

	}

	fmt.Printf("\nSearched %d positions!", searchCount)

	return bestMove

}

func Evaluate(board *gocheckers.CheckersBoard) int {
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

func AlphaBeta(board *gocheckers.CheckersBoard, depth int, maximizingPlayer bool) int {
	searchCount += 1
	if depth == 0 {
		return Evaluate(board)
	}

	moves := board.GenerateMoves() // Children nodes of this node

	if len(moves) == 0 {
		return Evaluate(board)
	}

	if maximizingPlayer {
		bestValue := math.MinInt // We are looking to maximize this value

		for _, move := range moves {
			boardCopy := board.GenerateDeepCopy()
			board.MakeMove(move)

			bestValue = Max(bestValue, AlphaBeta(boardCopy, depth-1, false))

			// When the best value is greater than beta (which is the current maximum value), then we can return from this search.
			// The opposing player will never pick this path since there are other options that are guarenteed to be better for them
			if bestValue > beta && depth < 6 {
				break
			}

			// Alpha is the best score that the maximizing player can get
			alpha = Max(alpha, bestValue)

		}

		return bestValue

	} else {

		bestValue := math.MaxInt // We are looking to minimize this value

		for _, move := range moves {
			boardCopy := board.GenerateDeepCopy()
			boardCopy.MakeMove(move)

			bestValue = Min(bestValue, AlphaBeta(board, depth-1, true))

			// If we find a value along this search path that is less than the alpha, we can stop searching. The other player will
			// never choose to go down this path since there are other options that put them at an advantage more than this move.
			if bestValue < alpha && depth < 6 {
				break
			}

			// This is tracking our best possible move.
			beta = Min(beta, bestValue)
		}

		return bestValue

	}

}

func main() {
	board := gocheckers.NewCheckersBoard()

	for {
		if board.GetTurn() == gocheckers.Black {
			fmt.Println("It's your move Black!")

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

			board.MakeMove(moves[move])

		} else {
			fmt.Println("It's your turn White!")
			fmt.Println(board.ToString())
			move := GetNextMove(board)
			fmt.Printf("White is making the move %v\n", move)
			board.MakeMove(move)
		}
	}
}

// Extra functions
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}

func MinOf[T Number]() (r T) {
	switch x := any(&r).(type) {
	case *int:
		*x = math.MinInt
	case *int8:
		*x = math.MinInt8
	case *int16:
		*x = math.MinInt16
	case *int32:
		*x = math.MinInt32
	case *int64:
		*x = math.MinInt64
	case *uint:
		*x = 0
	case *uint8:
		*x = 0
	case *uint16:
		*x = 0
	case *uint32:
		*x = 0
	case *uint64:
		*x = 0
	case *float32:
		*x = -math.MaxFloat32
	case *float64:
		*x = -math.MaxFloat64
	default:
		panic("unreachable")
	}
	return
}

func MaxOf[T Number]() (r T) {
	switch x := any(&r).(type) {
	case *int:
		*x = math.MaxInt
	case *int8:
		*x = math.MaxInt8
	case *int16:
		*x = math.MaxInt16
	case *int32:
		*x = math.MaxInt32
	case *int64:
		*x = math.MaxInt64
	case *uint:
		*x = math.MaxUint
	case *uint8:
		*x = math.MaxUint8
	case *uint16:
		*x = math.MaxUint16
	case *uint32:
		*x = math.MaxUint32
	case *uint64:
		*x = math.MaxUint64
	case *float32:
		*x = math.MaxFloat32
	case *float64:
		*x = math.MaxFloat64
	default:
		panic("unreachable")
	}
	return
}

func Max[T Number](n1 T, n2 T) (r T) {
	if n1 > n2 {
		return n1
	}

	return n2
}

func Min[T Number](n1 T, n2 T) (r T) {
	if n1 > n2 {
		return n2
	}

	return n1
}
