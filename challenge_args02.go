package main

import (
    "os"
    "fmt"
)

func main() {
    for x, y := range os.Args[1:] {
        fmt.Println(x+1, y)
    }
}
