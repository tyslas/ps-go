package main

import (
	"fmt"
)

func main() {

	//maps are unsafe for concurrency - reference type
	//Go doesn't always put these entries in the same order as declared
	testMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	for key, value := range testMap {
		fmt.Printf("\n key is: %v -- value is: %v", key, value)
	}

	fmt.Println("\n key a:", testMap["a"])
	testMap["b"] = 77
	testMap["f"] = 1886
	fmt.Println("\n updated map:", testMap)

	delete(testMap, "f")
	fmt.Println("\n updated map:", testMap)	

/* 	//declare map
	leagueTitles := make(map[string]int)
	//initialize values
	leagueTitles["sunderland"] = 6
	leagueTitles["newcastle"] = 4

	//declare and initialize (composite-literal form)
	recentHead2Head := map[string]int {
		"sunderland": 5,
		"newcastle": 0,
	}

	fmt.Printf("\nleague titles: %v\nrecent head to heads: %v\n", leagueTitles, recentHead2Head)
 */
}