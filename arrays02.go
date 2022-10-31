package main

import "fmt"

func main() {
    var myArray [3]string
    myArray[0] = "United States"
    myArray[1] = "Canada"
    myArray[2] = "Japan"


    fmt.Println(myArray)
    fmt.Println(myArray[0])
    fmt.Println(myArray[1])
    fmt.Println(myArray[2])
}
