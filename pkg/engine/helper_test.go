package main

import "testing"

func TestSq120To64(t *testing.T) {
	got := sq120To64(53)
	if got != 26 {
		t.Errorf("sq120To64(53) = %d; want 26", got)
	}
}

func TestSq64To120(t *testing.T) {
	got := sq64To120(26)
	if got != 53 {
		t.Errorf("sq64To120(26) = %d; want 53", got)
	}
}

func TestGetAlgebraic(t *testing.T) {
	got := getAlgebraic(55)
	if got != "e4" {
		t.Errorf("getAlgebraic(55) = %s; want e4", got)
	}
}

func TestGetCPermString(t *testing.T) {
	got := getCPermString(0)
	if got != "-" {
		t.Errorf("getCPermString(0) = %s; want -", got)
	}
	got = getCPermString(15)
	if got != "KQkq" {
		t.Errorf("getCPermString(15) = %s; want KQkq", got)
	}
	got = getCPermString(9)
	if got != "Kq" {
		t.Errorf("getCPermString(9) = %s; want Kq", got)
	}
}
