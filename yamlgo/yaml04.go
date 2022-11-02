package main

import (
    "fmt"
    "os"
    "log"
    "gopkg.in/yaml.v3"
)

type User struct {
    Name string
    Occupation string
}

func main() {
    users := map[string]User{"user 1": {"John Doe", "gardener"},
          "user 2": {"Lucy Black", "teacher"}}

    data, err := yaml.Marshal(&users)

    if err != nil {
        log.Fatal(err)
    }

    err2 := os.WriteFile("users.yaml", data, 0644)

    if err2 != nil {

       log.Fatal(err2)
    }

     fmt.Println("data written")
}