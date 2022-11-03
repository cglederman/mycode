package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)

func main() {
   resp, err := http.Get("http://example.com")
   // check for errors
    if err != nil {
        log.Fatal(err)
    }

    // the client must close the response body when finished
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))

}
