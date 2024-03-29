package main

import (
	"fmt"
)

func main() {
	bestFinish := bestLeagueFinishes(13, 10, 13, 17, 14, 16, 6, 10, 8)
	fmt.Println(bestFinish)
}

//elipses - defines a variadic funct that accepts n integers
func bestLeagueFinishes(finishes ...int) int {
	best := finishes[0]

	for _, i := range finishes {
		if i < best {
			best = i
		}
	}
	
	return best
}