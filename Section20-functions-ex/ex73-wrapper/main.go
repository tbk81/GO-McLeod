package main

import (
	"fmt"
	"time"
)

func main() {
	timefunc(doWork)
}

func doWork() {
	for i := range 2000 {
		fmt.Println(i)
	}
}

func timefunc(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

/*
Simple example to illustrate the concept of a wrapper function in Go
*/
