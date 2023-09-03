package main

import "testing"

func TestAdd(t *testing.T) {
	// arrange
	l, r := 1, 2
	expect := 3

	// act
	got := Add(l, r)

	// assert
	if got != expect {
		t.Errorf("got %v, want %v", got, expect)
	}
}
