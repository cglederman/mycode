package main

import (
    "fmt"
)

func main() {
    futurama := make(map[string]interface{})

    // now we can create a map of mixed types
    futurama["name"] = "Turanga Leela"  // string: string
    futurama["age"] = 30                // string: int
    futurama["height"] = 182.5          // string: float
    fmt.Printf("%+v\n", futurama)
    fmt.Printf("%+T\n", futurama["name"])
    fmt.Printf("%+T\n", futurama["age"])
    fmt.Printf("%+V\n", futurama)

}
