package main

import "fmt"

func main() {
    s := make([]string, 3)
    fmt.Println("slice s:", s)

    s[0] = "a"
    s[1] = "b"
    s[2] = "c"

    fmt.Println("slice s:", s)
    fmt.Println("get s[2]:", s[2])
    
    fmt.Println("len s:", len(s))

    s = append (s, "d")
    s = append(s, "e", "f")
    fmt.Println("append s:", s)
    c := make([]string, len(s))
    copy(c,s)
    fmt.Println("copy c:", c)
    
    c[2] = "updated"
    fmt.Println("original slice s:", s)
    fmt.Println("copy slice c:", c)

    l := s[2:5]
    fmt.Println("sl1:", l)

    l = s[:5]
    fmt.Println("sl2:", l)

    l = s[2:]
    fmt.Println("sl3:", l)
}
