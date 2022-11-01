package main

import "fmt"

func hello(names ...string) {
    fmt.Println(names, " ")
    total := 0

    for _, name := range names {
        total += 1
        fmt.Println("Hello", name)
    }
    fmt.Println(total)
}

func main() {
    hello("larry", "steve")
    hello("jane", "raj", "tom")
    names := []string{"larry", "raj", "harry", "beth"}
    hello(names...)
}

