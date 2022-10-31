package main

import (
    "fmt"
)

func rollchar(firstName, lastName string) (string, error) {
    return firstName + " the " + lastName, nil
}

func main() {
    fmt.Println("Welcome to the Character Generator")
    playerChar, err := rollchar("Gandalf", "Grey")
    if err != nil {
        fmt.Println("Error while spawning your requested charater")
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}
