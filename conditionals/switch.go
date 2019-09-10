package main

import (
	"fmt"
)

func main() {

	//types must match
	//fallthrough automatically executes the next consecutive case
	switch "linux" {
	case "linux":
		fmt.Println("\nhere are some recommended Linux courses")
		fallthrough
	case "docker":
		fmt.Println("\nhere are some recommended Docker courses")
	case "windows":
		fmt.Println("\nhere are some recommended Windows courses")
	default: 
		fmt.Println("\nno matches - please take a look at the most popular courses below")
	}
}