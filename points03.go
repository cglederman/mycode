package main

import (
    "fmt"
)

func noPointer() string {
    return "string"
}

func pointerTest() *string {
    return nil
}

func pointerTestTwo() *string {
    s := "string"
    return &s
}

func main() {
    fmt.Println(noPointer())
    fmt.Println(pointerTest())
    fmt.Println(pointerTestTwo())
    s := pointerTestTwo()
    fmt.Println(s)
    sp := *s
    fmt.Println(sp)
}
