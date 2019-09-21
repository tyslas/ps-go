package main

import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	goDur, _ := time.ParseDuration("10ms")
	runtime.GOMAXPROCS(2)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("hello")
			time.Sleep(goDur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(goDur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}