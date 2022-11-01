package main

import (
    "fmt"
    "time"
    "math/rand"
)

func main() {
     rand.Seed(time.Now().UnixNano())
     names := []string{"Joe", "Mary", "Jay", "Raj", "Tom", "Julia"}
     fmt.Println(names)
     if x := names[rand.Intn(len(names))]; x == "Joe" {
        fmt.Println("Joe works out of the NorthEast")
     } else if x == "Tom" {
        fmt.Println("Tom lives in the Pacific NorthWest")
     } else {
        fmt.Println("Hmm, I don't know much about the instructor,", x)
     }
}
