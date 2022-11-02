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
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
}
