package main

import (
    "fmt"
    "os/user"
)

type Userdata struct {
    na string
    un string
    hd string
    em string
}

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    currentUser := Userdata{user.Uid, user.Username, user.HomeDir, "rzfeeser@example.com"}
    fmt.Println("User id:", currentUser.na)
    fmt.Println("Username:", currentUser.un)
    fmt.Println("Home Directory:", currentUser.hd)
    fmt.Println("Email:", currentUser.em)

}
