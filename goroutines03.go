package main

import (
    "fmt"
    "sync"
)

var (
    counter int32
    wg sync.WaitGroup
    mutex sync.Mutex
)

func main() {
    wg.Add(4)

    go SuperPower("controlling elements", 2)
    go SuperPower("energy projection", 5)
    go SuperPower("flight", 6)
    go SuperPower("healing", 2)

    wg.Wait() // Wait for the goroutines to finish.
    fmt.Println("Total Mutant Powers Used:", counter)

}

func SuperPower(power string, num int) {
    defer wg.Done() // Schedule the call to Done to tell main we are done.

    for i := 0; i < num; i++ {
        mutex.Lock()    // a mutex describes a boundary past which only a single thread may operate at a time
        {
            fmt.Println(power)
            counter++
        }
        mutex.Unlock()  // exit the mutex
    }
}
