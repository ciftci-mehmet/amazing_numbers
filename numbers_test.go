package main

import "testing"

func TestIsNaturalNumber(t *testing.T) {
	tables := []struct {
		input string
		want  bool
	}{
		{"15", true},
		{"3.5", false},
		{"-2", false},
		{"qwe", false},
	}

	for _, table := range tables {
		got := isNaturalNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsEvenNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{5, false},
		{6, true},
		{13, false},
		{14, true},
		{17, false},
		{120, true},
	}

	for _, table := range tables {
		got := isEvenNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsOddNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{5, true},
		{6, false},
		{13, true},
		{14, false},
		{17, true},
		{120, false},
	}

	for _, table := range tables {
		got := isOddNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsBuzzNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{5, false},
		{7, true},
		{13, false},
		{14, true},
		{17, true},
		{19, false},
	}

	for _, table := range tables {
		got := isBuzzNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsDuckNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{5, false},
		{50, true},
		{142, false},
		{102, true},
	}

	for _, table := range tables {
		got := isDuckNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsPalindromicNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{5, true},
		{50, false},
		{101, true},
		{1234, false},
		{17371, true},
		{12388321, true},
	}

	for _, table := range tables {
		got := isPalindromicNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsGapfulNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{12, false},
		{123, false},
		{132, true},
		{7881, true},
	}

	for _, table := range tables {
		got := isGapfulNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}

func TestIsSpyNumber(t *testing.T) {
	tables := []struct {
		input int
		want  bool
	}{
		{1, true},
		{12, false},
		{22, true},
		{132, true},
		{325, false},
	}

	for _, table := range tables {
		got := isSpyNumber(table.input)
		if got != table.want {
			t.Errorf("Incorrect, \ninput: %v,\n want: %v, \n  got: %v", table.input, table.want, got)
		}
	}
}
