package main

import (
    "fmt"
    "github.com/cglederman/supermario/models"
)


func main() {
    mario := models.Player {3, 1, []string{"mushroom"}}
    mario.display()
    mario.Greenmushroom()
    mario.display()
    mario.Pickup("mushroom")
    mario.display()
    fmt.Println(mario.CanWhistle())
}

