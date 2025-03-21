package main

import "testing"

func TestAdd(t *testing.T) {
	total := add(5, 5)
	if total != 10 {
		t.Errorf("Add func was incorrect, got %d, want %d", total, 10)
	}
}

func TestSubtract(t *testing.T) {
	total := subtract(10, 5)
	if total != 5 {
		t.Errorf("Subtract func was incorrect, got %d, want %d", total, 5)
	}
}
