package main

import (
    "fmt"
    "sort"
)

type Person struct {
    Name string
    Age int
    Height float64
}


func (p Person) String() string {
    return fmt.Sprintf("%s: %d %.2f", p.Name, p.Age, p.Height)
}

type ByAge []Person
type ByHeight []Person

func (a ByAge) Len() int {
    return len(a)
}

func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

func (a ByHeight) Len() int {
    return len(a)
}

func (a ByHeight) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a ByHeight) Less(i, j int) bool {
    return a[i].Height < a[j].Height
}

func main() {
    people := []Person{
        {"Bob", 31, 68*2.54},
        {"John", 42, 72*2.54},
        {"Michael", 17,62*2.54},
        {"Jenny", 26, 55*2.54},
    }
    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)
    sort.Sort(ByHeight(people))
    fmt.Println(people)

}

    
