package main

import "fmt"

func main() {
    array := [5]int{1,2,3,4,5}
    slice := array[:]
    slice[1] = 7

    fmt.Println("Modifiying slice")
    fmt.Println("Array =", array)
    fmt.Println("Slice =", slice)
    
    array[1] = 2
    fmt.Println("Modifiing underlying array")
    fmt.Println("Array = ", array)
    fmt.Println("Slice = ", slice)


}
