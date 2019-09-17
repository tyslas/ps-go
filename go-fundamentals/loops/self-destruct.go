package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\ninitiating count down...")
	for timer := 10; timer >= 0; timer-- {

		// if timer == 0 {
		// 	fmt.Println("\nBOOM!")
		// 	break
		// }
		if timer % 2 == 0 {
			continue
		}

		fmt.Println("\n", timer)
		time.Sleep(1 * time.Second)
	}
}