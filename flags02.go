package main

import (
    "flag"
    "fmt"
)

var (
    env *string
    port *int
)

func main() {
    env := flag.String("env", "development", "current environment")
    port := flag.Int("port", 3000, "port number")
    flag.Parse()

    fmt.Println("The value of env:", *env)
    fmt.Println("The port has been set to:", *port)

}
