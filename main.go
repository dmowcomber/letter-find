package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

func main() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Key press => " + string(r))
	}
}
