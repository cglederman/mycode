package main

import ("fmt"
        "os")


var user = os.Getenv("USER")
var city = "Horsham"
var state = "PA"

func main() {
    fmt.Println("user name", user)
    fmt.Println("city:", city," state: ", state)
}
