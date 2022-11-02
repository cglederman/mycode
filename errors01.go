package main

import (
    "errors"
    "fmt"
)

func main() {
    err := errors.New("we don't have the power") // returns a new error with a static message
    fmt.Println("Scotty says:", err)

}
