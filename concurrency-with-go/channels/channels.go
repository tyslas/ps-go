package main

import (
	"fmt"
)

func main() {
	//instantiate a channel that accepts strings with a buffer of 1
	ch := make(chan string, 1)
	
	ch <- "hello"

	fmt.Println(<-ch)
}