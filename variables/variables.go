package main

import (
    "fmt"
    "reflect"
)

func main() {
    name := "Tito"               //name of subscriber
    course := "Go Fundamentals"  //course being viewed

    fmt.Println("\nHi", name, "you're currently watching", course)

    reference := &course
    fmt.Println("\nreference var is type ", reflect.TypeOf(reference))
    changeCourse(reference)
    fmt.Println("\nyou are now watching course ", course)
}
  
func changeCourse(course *string) string {
    //not a "new" variable declaration - don't need ":=" 
    *course = "Docker Deep Dive"
    fmt.Println("\ntrying to change your course to ", *course)
    return *course
}