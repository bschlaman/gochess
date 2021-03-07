package main

import "testing"

func TestParseFEN(t *testing.T) {
	inputFEN = ""
	bs := NewBoardState()
	got := ParseFEN(inputFEN, bs)
	if got != 26 {
		t.Errorf("sq120To64(53) = %d; want 26", got)
	}
}
