package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{"docker deep dive", "docker clustering", "docker & kubernetes"}
	coursesCompleted := []string{"docker deep dive", "go fundamentals", "intro to docker"}

	for _, i := range coursesInProg {
		// fmt.Println(i) 
		if timer % 2 == 0 {
			continue
		}
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println("ERROR:", j, "is in both lists")
			}
		}
	}
}