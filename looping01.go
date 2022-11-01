package main

import "fmt"

func main() {
    total := 0

    for alta := 0; alta < 4; alta++ {
        fmt.Println("The value of alta -> ", alta)
        
        total += alta   // use total to keep track of alta
    }

fmt.Println("\nThe value of total is", total)
    
}
