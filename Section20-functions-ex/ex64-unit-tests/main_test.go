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

func TestDoMath(t *testing.T) {
	total := doMath(5, 5, add)
	if total != 10 {
		t.Errorf("Do math func was incorrect, got %d, want %d", total, 10)
	}

}
func TestDoMathAgain(t *testing.T) {
	total := doMath(10, 5, subtract)
	if total != 5 {
		t.Errorf("Do math func was incorrect, got %d, want %d", total, 5)
	}
}
