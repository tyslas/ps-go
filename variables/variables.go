package main

import (
    "fmt"
    "reflect"
)

func main() {
    name := "Tito"               //name of subscriber
    course := "Go Fundamentals"  //course being viewed
    module := 3.2                //current position in course
    ptr := &module               //reference to memory addres

    fmt.Println("Name is set to ", name, " and is type ", reflect.TypeOf(name))
    fmt.Println("Course is set to ", course, " and is type ", reflect.TypeOf(course))
    fmt.Println("Module is set to", module, " and is type ", reflect.TypeOf(module))
    fmt.Println("Memory address of *module* variable is ", &module)
    fmt.Println("my ptr variable value is ", ptr, " de-referencing ptr gives us ", *ptr)
}
 