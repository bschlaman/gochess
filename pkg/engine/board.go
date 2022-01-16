package main

import (
	"fmt"
	"strconv"
)

type Side int
type Type int

const (
	White = iota
	Black
	Knight Type = iota
	Bishop
	Rook
	Queen
	King
	PieceChars  string = ".PNBRQKpnbrqkx"
	CastleChars string = "KQkq"
	//NumDirections []int = {8,4,4,8,8}
)

func (s Side) String() string {
	return [...]string{"white", "black"}[s]
}

type Move struct {
	To    int
	From  int
	Flags int
}

type Irrev struct {
	move          Move
	EnPasSq       int
	CPerm         int
	CapturedPiece int
	Pinned        uint64
}

type BoardState struct {
	Board      [120]int
	Ply        int
	SideToMove Side
	EnPasSq    int
	// 0 0 0 0
	CPerm   int
	KingSq  [2]int
	Pinned  uint64
	History [200]Irrev
}

func NewBoardState() *BoardState {
	var BS BoardState
	return &BS
}

const (
	NORMAL_MODE = iota
	FEN_MODE
	PERFT_MODE
	O_64_BOARD
	O_BOARD_STATE
	O_120_BOARD
	O_PINNED
)

// helper functions
func sq120To64(sq120 int) int {
	return sq120 - 17 - 2*(sq120-sq120%10)/10
}

func sq64To120(sq64 int) int {
	return sq64 + 21 + 2*(sq64-sq64%8)/8
}

func getAlgebraic(sq120 int) string {
	sq64 := sq120To64(sq120)
	file := (sq64 % 8) + 'a'
	rank := 1 + ((sq64 - sq64%8) / 8)
	return string(rune(file)) + strconv.Itoa(rank)
}

func getCPermString(cperm int) string {
	if cperm == 0 {
		return "-"
	} else {
		var cpermString string
		for i := 0; i < 4; i++ {
			if 1<<(3-i)&cperm > 0 {
				cpermString += string(CastleChars[i])
			}
		}
		return cpermString
	}
}

func (bs *BoardState) PrintBoard(opt int) {
	switch opt {
	case O_64_BOARD:
		fmt.Println(colorize(YEL, " ---- Game Board ---- "))
		for i := 0; i < 64; i++ {
			if i%8 == 0 {
				fmt.Printf(" %d ", 8-(i-i%8)/8)
			}
			index := 56 + i - 2*(i-i%8)
			fmt.Printf("%c ", PieceChars[bs.Board[index]])
			if (i+1)%8 == 0 {
				fmt.Println()
			}
		}
		fmt.Printf("   ")
		for i := 0; i < 8; i++ {
			fmt.Printf("%c ", rune(i+'a'))
		}
		fmt.Printf("\n\n")
	case O_BOARD_STATE:
		fmt.Println(colorize(YEL, " ---- Board State ---- "))
		fmt.Println(colorize(BLU, "ply:"), bs.Ply)
		fmt.Println(colorize(BLU, "side to move:"), bs.SideToMove.String())
		fmt.Println(colorize(BLU, "en passant:"), getAlgebraic(bs.EnPasSq))
		fmt.Println(colorize(BLU, "castle perms:"), getCPermString(bs.EnPasSq))
		fmt.Println(colorize(BLU, "white in check:"), "")
		fmt.Println(colorize(BLU, "black in check:"), "")
		fmt.Println(colorize(BLU, "white king sq:"), getAlgebraic(bs.KingSq[White]))
		fmt.Println(colorize(BLU, "black king sq:"), getAlgebraic(bs.KingSq[Black]))
		fmt.Println(colorize(BLU, "eval of position for"), bs.SideToMove.String(), ":")
	case O_120_BOARD:
		fmt.Println(colorize(YEL, " ---- 120 Board ---- "))
	case O_PINNED:
		fmt.Println(colorize(YEL, " ---- Pinned Pieces ---- "))
	default:
		fmt.Println(colorize(RED, "Error: invalid option"))
	}
}
