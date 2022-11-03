package main

import (
    "fmt"
)

type Wizard struct{}

func (w Wizard) Defend() string {
    return "Expelliarmus"
}

// Wizard Receiver Function (method) - A Wizard has the unique ability to cast a Forget() spell
func (w Wizard) Forget() string {
    return "Obliviate"
}

type Barbarian struct{}

func (b Barbarian) Defend() string {
    return "Dodge"
}

func main() {

    gandalf := Wizard{}
    fmt.Println("Wizard defends:", gandalf.Defend())
    
    // repeat the same pattern for the Barbarian
    conan := Barbarian{}
    fmt.Println("Barbarian defends:", conan.Defend())

    // continue the pattern for any other players


    fmt.Println("Wizard makes us all forget:", gandalf.Forget())

}
