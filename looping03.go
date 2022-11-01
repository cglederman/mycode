package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    drive := 0
    fmt.Print("Get a mulligan on any drive under 60 yards.\n")
    rand.Seed(time.Now().UnixNano())
    // typically after "for" is the init var, which is omitted
    // no post statement appears after the last semi-colon
    for ; drive <= 60; {           // same as: for drive <= 60
        drive = rand.Intn(251)
        fmt.Print("SWING!\n")
    }
    fmt.Println("Your longest drive was", drive, "yards")
}
