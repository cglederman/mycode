package main

import (
    "fmt"
    "github.com/cglederman/supermario/models"
)


func main() {
    mario := models.Player {3, 1, []string{"mushroom"}}
    mario.Display()
    mario.Greenmushroom()
    mario.Display()
    mario.Pickup("mushroom")
    mario.Display()
    fmt.Println(mario.CanWhistle())
}

