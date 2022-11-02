package main

import (
    "fmt"
    "io"
    "os"
    "log"
    "gopkg.in/yaml.v3"
)

type User struct {
    Name string
    Occupation string
}

func main() {
     fs, err := os.Open("villains.yaml")
     if err != nil {
          log.Fatal(err)   // provides standard datetime formatting and then calls os.exit(1)
     }

     // read the contents of, "villains.yaml"
     yfile, err := io.ReadAll(fs)

     // error checking
     if err != nil {
          log.Fatal(err)   // provides standard datetime formatting and then calls os.exit(1)
     }

     data := make(map[string]User)

     err2 := yaml.Unmarshal(yfile, &data)

     if err2 != nil {
         log.Fatal(err2)
     }

     for k, v := range data {
         fmt.Printf("%s: %+V\n", k, v)
     }
 }
