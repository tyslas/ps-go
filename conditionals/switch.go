package main

import (
	"fmt"
)

func main() {

	//types must match
	//only a single case statement will be executed per switch block
	switch "aws" {
	case "linux":
		fmt.Println("\nhere are some recommended Linux courses")
	case "docker":
		fmt.Println("\nhere are some recommended Docker courses")
	case "windows":
		fmt.Println("\nhere are some recommended Windows courses")
	default: 
		fmt.Println("\nno matches - please take a look at the most popular courses below")
	}
}