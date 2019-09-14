package main

import (
	"fmt"
)

func main() {

	//define custom type
	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	//initializes with "zero" value
	// var DockerDeepDive courseMeta
	//creates a pointer & initializes with "zero" value
	// DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Nigel P",
		Level: "Intermediate",
		Rating: 5,
	}

	fmt.Println("\ncourse author:", DockerDeepDive.Author)
}