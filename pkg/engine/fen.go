package main

import "fmt"

func ParseFEN(fen string, bs BoardState) {
	for i, char := range fen {
		numEmpty := 1
		switch char {
		case "p":
			piece := bP
			break
		case "r":
			piece := bR
			break
		case "n":
			piece := bN
			break
		case "b":
			piece := bB
			break
		case "q":
			piece := bQ
			break
		case "k":
			piece := bK
			break
		case "P":
			piece := wP
			break
		case "R":
			piece := wR
			break
		case "N":
			piece := wN
			break
		case "B":
			piece := wB
			break
		case "Q":
			piece := wQ
			break
		case "K":
			piece := wK
			break

		case "1":
		case "2":
		case "3":
		case "4":
		case "5":
		case "6":
		case "7":
		case "8":
			piece = EMPTY
			num = *fen - "0"
			break

		case "/":
		case " ":
			rank--
			file = 0
			fen++
			continue

		default:
			fmt.Println(colorize(RED, "Error with FEN"))
			return -1
		}
	}
}
