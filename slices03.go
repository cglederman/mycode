package main

import "fmt"

func main() {
    numbers := make([]int, 3, 5)
    fmt.Println("numbers = ", numbers)
    fmt.Println("length = ", len(numbers))
    fmt.Println("capacity = ", cap(numbers))

    numbers = numbers[0:5]

    numbers[4] = 5

    fmt.Println("Inceasing length from 3 to 5")
    fmt.Println("numbers =", numbers)
    fmt.Println("length =", len(numbers))
    fmt.Println("capcity =", cap(numbers))

    numbers = append(numbers, 66, 777, 8888)
    fmt.Println("numbers =", numbers)
    fmt.Println("capcity =", cap(numbers))
}

