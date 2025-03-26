package main

import (
	"log"
	"testing"
)

func TestSaySomething(t *testing.T) {
	got := saySomething("Hellooooo gopherman")
	want := "This is something I said to print, Hellooooo gopherman"
	if got != want {
		log.Fatalf("Error - got: %s\nwant: %s", got, want)
	}
}
