package main

import (
    "fmt"
    "reflect"
)

func main() {
    name := "Tito"               //name of subscriber
    course := "Go Fundamentals"  //course being viewed
    module := 3.2                //current position in course

    fmt.Println("Name is set to ", name, " and is type ", reflect.TypeOf(name))
    fmt.Println("Course is set to ", course, " and is type ", reflect.TypeOf(course))
    fmt.Println("Module is set to", module, " and is type ", reflect.TypeOf(module))
}
 