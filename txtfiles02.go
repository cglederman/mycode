package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("zelda.txt")

    if err != nil {
        log.Fatalf("failed opening file: %s", err)
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string 
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    for _, eachline := range lines {
        fmt.Println(eachline)
    }
    
}
