package main

import (
	"flag"
	"fmt"
	"os"
)

const StartFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var inputFEN string

func parseArgs() {
	fen := flag.String("f", StartFEN, "fen notation of position")
	flag.Parse()
	inputFEN = *fen
	args := os.Args[1:]
	fmt.Println(args)
}

func main() {
	fmt.Println("Chess!")
	parseArgs()
	fmt.Println("input", inputFEN)

	bs := NewBoardState()
	bs.PrintBoard(O_64_BOARD)
	bs.PrintBoard(O_BOARD_STATE)
}
