package main

import (
    "fmt"
)


func main() {
    var q int = 42
    var p *int
    p = &q
    fmt.Println(p)
}
