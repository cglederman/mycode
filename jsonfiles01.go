package main

import (
    "encoding/json"
    "fmt"
    "io"
    "os"
    "strconv"
)

type Users struct {
    Users []User `json:"users"`
}

type User struct {
    Name string `json:"name"`
    Type string `json:"type:"`
    Age  int    `json:"Age"`
    Social Social `json:"social"`
}

type Social struct {
    Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main() {
    jsonFile, err := os.Open("users.json")

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully Opened users.json")
    defer jsonFile.Close()

    byteValue, _ := io.ReadAll(jsonFile)
    var users Users

    json.Unmarshal(byteValue, &users)

    for i := 0; i < len(users.Users); i++ {
        fmt.Println("User Type: " + users.Users[i].Type)
        fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        fmt.Println("User Name: " + users.Users[i].Name)
        fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }
}
