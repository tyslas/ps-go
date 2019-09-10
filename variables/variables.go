package main

import (
    "fmt"
)

func main() {
    name := "Tito"               //name of subscriber
    course := "Go Fundamentals"  //course being viewed

    fmt.Println("\nHi", name, "you're currently watching", course)

    changeCourse(course)
    fmt.Println("\nyou are now watching course ", course)
}
 
func changeCourse(course string) string {
    //not a "new" variable declaration - don't need ":=" 
    course = "Docker Deep Dive"
    fmt.Println("\ntrying to change your course to ", course)
    return course
}