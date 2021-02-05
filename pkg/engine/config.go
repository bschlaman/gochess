package main

type Color string

const (
	BLK Color = "\u001b[30m"
	RED       = "\u001b[31m"
	GRN       = "\u001b[32m"
	YEL       = "\u001b[33m"
	BLU       = "\u001b[34m"
	RES       = "\u001b[0m"
)

func colorize(color Color, msg string) string {
	return string(color) + msg + string(RES)
}
