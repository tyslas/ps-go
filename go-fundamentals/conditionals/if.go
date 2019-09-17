package main

import (
	"fmt"
)

func main() {
	//variables to store AP top 25 rankings
	usc := 24
	texas := 12

	if usc < texas {
		fmt.Println("wow, USC really found a way to win games and climb the rankings")
	} else if usc > texas {
		fmt.Println("\nUSC has a worse ranking than UT")
	} else {
		fmt.Println("somehow the AP couldn't decide on which team has the upper hand")
	}
}