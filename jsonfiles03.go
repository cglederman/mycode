package main

import (
    "encoding/json"
    "os"
)

func main() {
    type Person struct {
        Fn string
        Ln string
    }

    type ColorGroup struct {
        ID     int      // uppercase - "OK" at function level because it is a field name
        name   string   // lowercase
        colors []string // lowercase
        p      Person `json:"Person"`   // lowercase (struct, not JSON tag)
    }

    per := Person{Fn: "RZ",
        Ln: "Feeser",
    }

    // group is the TYPE ColorGroup
    group := ColorGroup{
        ID:     24601,
        name:   "Greens",
        colors: []string{"Moss", "Shamrock", "Lime", "Hunter"},
        p:      per,
    }

    // serialize values from struct to json format
    // our struct "group" is defined in the main package
    b, err := json.Marshal(group) // requesting interaction between json package and main package
    // result is ONLY the "exported" (capitalized) names will be found!

    if err != nil {
        panic(err)
    }

    // print to standard out
    os.Stdout.Write(b) // the "ONLY" thing that the json package could find was ColorGroup.ID!
}
