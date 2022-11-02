package main

import (
    "errors"
    "fmt"
)

func endit() error {
    return errors.New("we don't have the power")
}

func main() {
    err := endit()

    if err != nil {
        fmt.Println("We detected and error:", err)
        return  // return from main()
    }

    fmt.Println("Release the hounds!")
}
