package main 

import "fmt"


func main() {
    const uri = "https://example.org:6001/v2/snacks?"
    var r, q, s string = "req=snickers", "quantity=1", "size=king"

    url_text := fmt.Sprintf("%s%s&%s&%s", uri, r, q, s)
    fmt.Println(url_text) // finished URI
}
