package main

import (
	"fmt"
	"strings"
)

func main() {
	phrase := "these are the times that try men's souls.\n"

	words := strings.Split(phrase, " ")

	ch := make(chan string, len(words))

	for _, word := range words {
		ch <- word
	}

	close(ch)

	//using the for range loop will only pull messages while the channel is open
	//more succint than using a infinite for loop with an if\else statement 
	//to check whether the channel is closed
	for msg := range ch {
		fmt.Print(msg + " ")
	}
}