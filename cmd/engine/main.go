package main

import (
	"flag"
	"fmt"
	"os"

	"board"
)

const StartFEN = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

var inputFEN string

func printBoard() {
	fmt.Println(" ---- Game Board ---- ")
}

func parseArgs() {
	fen := flag.String("f", StartFEN, "fen notation of position")
	flag.Parse()
	inputFEN = *fen
	args := os.Args[1:]
	fmt.Println(args)
}

func main() {
	printBoard()
	fmt.Println("Chess!")
	parseArgs()
	fmt.Println("input", inputFEN)

	bs := board.BoardState
	fmt.Println(bs)
}
