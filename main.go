package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

func main() {
	// need at least 2 items, including the name of the program
	if len(os.Args) < 3 {
		fmt.Println("Not enough to pick from")
		os.Exit(1)
	}

	// everything except name of program
	options := os.Args[1:]

	// seed our random picking and pick
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	// pick
	fmt.Printf("%s\n", options[r.Intn(len(options))])
}