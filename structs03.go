package main

import (
    "fmt"
)

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{"Bob", 20})
    fmt.Println(person{name:"Alice", age:30})
    fmt.Println(person{name:"Fred"})

    s := person{age: 58, name: "Okon"}

    fmt.Println(s.name)
    fmt.Println(s.age)

    sp := &s
    fmt.Println(sp.age)

    sp.age = 51
    fmt.Println(sp.age)
    fmt.Println(s.age)

}
