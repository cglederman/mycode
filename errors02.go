package main

import (
    "fmt"
    "time"
)

func main() {
    err := fmt.Errorf("error occurred at: %v", time.Now()) // use a formatting verb to add info to the error
    fmt.Println("An error happened:", err)
}
