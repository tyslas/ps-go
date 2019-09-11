package main

import (
	"fmt"
)

func main() {

	mySlice := make([]int, 1, 4)
	fmt.Printf("length is: %d -- capacity is: %d\n", len(mySlice), cap(mySlice))

	for i := 1; i <17; i++ {
		mySlice = append(mySlice, i)
		if i == 16 {
			fmt.Printf("\ncapacity is: %d\n", cap(mySlice))
			break
		}
		fmt.Printf("\ncapacity is: %d", cap(mySlice))
	}
	fmt.Println("mySlice:", mySlice)

	slice4 := mySlice[1:6]
	fmt.Println("\nslice4 from first 5 elements of myslice:\n", slice4)
	
	// fmt.Println("\nassignment loop")
	assignmentValue := 11
	for i := 0; i < len(slice4); i++ {
		slice4[i] = assignmentValue
		assignmentValue = assignmentValue + 1
		fmt.Println("assignmentValue:", assignmentValue)
	}
	
	fmt.Println("\nrange loop")
	for _, i := range slice4 {
		fmt.Println("for range loop:", i)
	}

	//append a slice to a slice
	mySlice = append(mySlice, slice4...)
	fmt.Println("both slices:\n", mySlice)
	
	//declares a slice called myCourses - generic
	// myCourses := make([]string, 5, 10)
	
	//initialize w/values - capacity automagically set to length
	// myCourses := []string{"Kubernetes", "docker", "nsm"}

	// fmt.Printf("\nlength is: %d.\ncapacity is: %d\n", len(myCourses), cap(myCourses))

/* 	mySlice := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println("printing element 4:\n", mySlice[4]) */

	//access value at index i
/* 	mySlice[1] = 0
	fmt.Println("printing slice:\n", mySlice) */

	//intialize smaller slice of original
/* 	sliceOfSlice := mySlice[2:5]
	fmt.Println("printing sliceOfSlice:\n", sliceOfSlice) */
}