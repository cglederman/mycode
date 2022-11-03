package main

import (
    "fmt"
    "reflect"
)

type Wizard struct{}

func (w Wizard) Defend() string { 
    return "Expelliarmus" 
}

func (w Wizard) Attack() string {
    return "Zap"
}

// Wizard Receiver Function (method) - A Wizard has the unique ability to cast a Forget() spell
func (w Wizard) Forget() string {
    return "Obliviate"
}

type Barbarian struct{}

// Barbarian Receiver Function (method) - This is how we want the Barbarian to Defend()
func (b Barbarian) Defend() string {
    return "Dodge"
}
func (b Barbarian) Attack() string {
    return "Whack"
}

type Player interface {
    Defend() string
    Attack() string
}

func main() {
    gandalf := Wizard{}
    // fmt.Println("Wizard defends:", gandalf.Defend())
    conan := Barbarian{}
    // fmt.Println("Barbarian defends:", conan.Defend())

    // Thanks to the inteface Player
    // we have a common type-- it isn't the structure-- but what they DO
    // this allows us to create a slice of things that have the qualities of a "Player"
    players := []Player{gandalf, conan}

      for _, a := range players {
        fmt.Println(reflect.TypeOf(a).Name(), "defends:", a.Defend())
        fmt.Println(reflect.TypeOf(a).Name(), "attack:", a.Attack())
    }

    fmt.Println("Wizard makes us all forget:", gandalf.Forget())

}
