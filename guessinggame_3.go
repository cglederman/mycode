package main

import (
    "fmt"
    "math/rand"
    "time"
    "flag"
)

var (
    numGuesses *int
)
func main() {
    numGuesses := flag.Int("guess", 3, "number of guesses")
    flag.Parse()
    var secretNumber int
    fmt.Printf("You have %d tries ", *numGuesses)
    source := rand.NewSource(time.Now().UnixNano())
    randomizer := rand.New(source)

    if *numGuesses <= 5 {
        fmt.Println("Game: Guess a number between 0 and 10")
        secretNumber = randomizer.Intn(11)
    } else {
        fmt.Println("Game: Guess a number between 10 and 100")
        secretNumber = randomizer.Intn(91) + 10
        fmt.Println(secretNumber)
    }



    var guess int

    for try := 1; try <= *numGuesses; try++ {
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

       if try == *numGuesses {
            // if the number of tries is equal to 3, print game over and also the correct number

            fmt.Printf("Game over!!\n ")
            fmt.Printf("The correct number is %d\n", secretNumber)
        }

    }

}
