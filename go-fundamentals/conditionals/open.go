package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/path/to/test.txt")
	if err != nil {
		fmt.Println("\nerror returned was:", err)
	}
}