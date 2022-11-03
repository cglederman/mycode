package main

import (
    "fmt"
    "sync"
    "time"
) 

func worker(linkChan chan string, wg *sync.WaitGroup) {
    defer wg.Done()
    for url := range linkChan {
        time.Sleep(1 * time.Second)
        fmt.Printf("Done processing link #%s\n", url)
    }
}

func main() {
    yourLinksSlice := make([]string, 50)

    for i := 0; i < 50; i++ {
        yourLinksSlice[i] = fmt.Sprintf("%d", i+1)
    }
    fmt.Println(yourLinksSlice)

    lCh := make(chan string)

    wg := new(sync.WaitGroup)

    fmt.Println("Setting up recivers")
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go worker(lCh, wg)
    }

    fmt.Println("About to send data")
    for _, link := range yourLinksSlice {
        lCh <- link
    }

    close(lCh)
    wg.Wait()
    
}
