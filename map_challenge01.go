package main

import "fmt"


func main() {
    m := map[string]string{"Python":".py", 
                           "Go":".go", 
                            "Java":".java",
                            "Ansible":"yml",
                            "c++":".cpp",}
    fmt.Println(m)
    delete(m, "c++")
    fmt.Println(m)
    m["Julia"] = ".jl"
    fmt.Println(m)
}
