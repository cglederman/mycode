package main

import "fmt"

func main() {
    var gamescores = make(map[string]int)
    gamescores["Zerg"] = 9092

    fmt.Println("gamescoes:", gamescores)

    //var totalscore map[string]int
    //totalscore["minecraft"] = 912

    var network = make(map[string]int)
    network["Cisco"] = 190
    network["Arista"] = 56
    network["NetGear"] = 302

    fmt.Println("Network Device:", network)

    values := map[string]int{
        "abc": 123,
        "def": 345,
        "ghi": 567,
        "jkl": 897,
    }

    fmt.Println("values:", values)


}

