package main

import (
    "fmt"
)

type Astro struct {
    name string
    age int
    mission string
    isneeded bool
}

func main() {
    my_astros := []Astro{{"a1", 21, "m1", false}, 
                         {"a2", 22, "m2", true}, 
                         {"a3", 23, "m3", false}}

    fmt.Println(my_astros)
}
