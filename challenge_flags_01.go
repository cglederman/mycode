package main

import (
    "fmt"
    "flag"
    "io"
    "os"
    "bufio"
)
var (
    //filename = *string
    numLines *int
)
func main() {
    //filename := flag.String("filename", "flags03.go", "a filename")
    numLines := flag.Int("lines", 5, "number of lines")
    flag.Parse()
    var in io.Reader
    if filename := flag.Arg(0); filename != "" {
        f, err := os.Open(filename)

        if err != nil {
            fmt.Println("error opening file: err:", err)
            os.Exit(1)
        }

        defer f.Close()
        in = f
    } else {
        in = os.Stdin
    }

    buf := bufio.NewScanner(in)

    for i := 0; i < *numLines; i++ {
        if !buf.Scan() {
            break
        }
        fmt.Println(buf.Text())
    }
    if err := buf.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading: err:", err)
    }
}
