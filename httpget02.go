package main

import (
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    resp, err := http.Get("http://example.com")

    if err != nil {
      log.Fatal(err)
    }

    defer resp.Body.Close()

    _, err = io.Copy(os.Stdout, resp.Body)

   if err != nil {
      log.Fatal(err)
    }
}
