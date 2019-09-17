package main

import (
    "fmt"
    "os"
)

func main() {
    name := os.Getenv("USER")               //name of subscriber
    course := "Go Fundamentals"  //course being viewed

    fmt.Println("\nHi", name, "you're currently watching", course)

    reference := &course
    changeCourse(reference)
    fmt.Println("\nyou are now watching course ", course)
}
  
func changeCourse(course *string) string {
    //not a "new" variable declaration - don't need ":=" 
    *course = "Docker Deep Dive"
    fmt.Println("\ntrying to change your course to ", *course)
    return *course
}