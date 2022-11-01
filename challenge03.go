package main

import (
    "fmt"
)


type Player struct {
    Lives int
    Stage int
    Inventory []string
}

func (a *Player) Greenmushroom() {
    a.Lives++

}

func (a *Player) Pickup(item string) {
    a.Inventory = append(a.Inventory, item)
} 

func (a Player) CanWhistle() bool {
   return a.Stage >= 5 
} 

func (a Player) display() {
    fmt.Println("Lives: ", a.Lives)
    fmt.Println("Stage: ", a.Stage)
    fmt.Println("Inventory: ", a.Inventory)
}

func main() {
    mario := Player {3, 1, []string{"mushroom"}}
    mario.display()
    mario.Greenmushroom()
    mario.display()
    mario.Pickup("mushroom")
    mario.display()
    fmt.Println(mario.CanWhistle())
}

