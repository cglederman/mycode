package models

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

func (a Player) Display() {
    fmt.Println("Lives: ", a.Lives)
    fmt.Println("Stage: ", a.Stage)
    fmt.Println("Inventory: ", a.Inventory)
}
