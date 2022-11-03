package main

import (
    "fmt"
    "net/http"
    "io"
)

func main() {
  // define our URL (API) as a string
  url := "https://api.spacexdata.com/v3/capsules"
  // the HTTP method we wish to send (HTTP GET)
  method := "GET"

  // For control over HTTP client headers,
  // redirect policy, and other settings,
  // create a Client
  // A Client is an HTTP client
  client := &http.Client {
  }

  req, err  := http.NewRequest(method, url, nil)

   if err != nil {
    fmt.Println(err)
    return
  }

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()  // even if we have a problem, we still want this to happen
  body, err := io.ReadAll(res.Body)
   // error checking
  if err != nil {
    fmt.Println(err)
    return
  }

  //
  fmt.Println(string(body))
}
