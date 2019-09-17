package main

import (
	"fmt"
	"time"
	"sync"
	"runtime"
)

func main() {

	//set the number of threads that the go container has access to
	//can use for parrallelism
	runtime.GOMAXPROCS(2)

	//instruct main not to proceed beyond line 26
	//until both go-routines are complete
	var waitGrp sync.WaitGroup
	waitGrp.Add(2)

	go func() {
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("hello")
	}() //syntax for self-executing function

	go func() {
		defer waitGrp.Done()
		fmt.Println("pluralsight")
	}()

	waitGrp.Wait()
}