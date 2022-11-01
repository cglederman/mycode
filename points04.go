package main

import (
    "fmt"
)

type User struct {
    Name string
    Pets []string
}

func (u User) newPet() {
    u.Pets = append(u.Pets, "Lucy")
    fmt.Println(u)
}

func main() {
    u := User{Name: "Anna", Pets: []string{"Bailey"}}
    u.newPet()
    fmt.Println(u)
}
