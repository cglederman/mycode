package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    fmt.Println("Game: Guess a number between 0 and 10")
    fmt.Println("You have three(3) tries ")
    source := rand.NewSource(time.Now().UnixNano())

    randomizer := rand.New(source)
    secretNumber := randomizer.Intn(11)

    var guess int

    for try := 1; try <= 3; try++ {
        fmt.Println("Round:", try)
        fmt.Scan(&guess)

        if guess < secretNumber {
            fmt.Printf("Sorry, wrong guess ; number is too small\n ")
        } else if guess > secretNumber {
            fmt.Printf("Sorry, wrong guess ; number is too large\n ")
        } else {
            fmt.Printf("You win!\n")
            break
            // Print out "you win" message when the player guesses the correct number
        }

       if try == 3 {
            // if the number of tries is equal to 3, print game over and also the correct number

            fmt.Printf("Game over!!\n ")
            fmt.Printf("The correct number is %d\n", secretNumber)
        }

    }

}
