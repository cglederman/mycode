package main

import (
    "fmt"
    "io"
    "log"
    "os"

    "gopkg.in/yaml.v3"
)

func main() {
    fs, err := os.Open("startrek.yaml")
     if err != nil {
          log.Fatal(err)     // provides standard datetime formatting and then calls os.exit(1)
     }
    yfile, err := io.ReadAll(fs)
    if err != nil {
          log.Fatal(err)     // provides standard datetime formatting and then calls os.exit(1)
    }

    data := make(map[any]any)

    err2 := yaml.Unmarshal(yfile, &data)

    if err2 != nil {

          log.Fatal(err2)    // provides standard datetime formatting and then calls os.exit(1)
    }
    for k, v := range data {
        fmt.Printf("%s -> %d\n", k, v)
    }

}
