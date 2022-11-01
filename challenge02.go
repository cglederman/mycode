package main

import "fmt"

type Virtmach struct {
    ip string
    hostname string
    diskgb int
    ram int
}

func (a *Virtmach) increaseRam() {
    a.ram++
}

func (a *Virtmach) increaseDisk(gb int) {
    a.diskgb += gb
}

func (v Virtmach)display(){
    fmt.Println("Primary IP Address:", v.ip)    // primary IP address
    fmt.Println("Hostname:", v.hostname)        // hostname
    fmt.Println("Disk GB:", v.diskgb)           // diskgb
    fmt.Println("RAM:", v.ram)                  // ram
}

func main() {
    vm1 := Virtmach{"10.0.0.5", "zebra", 20, 8} 
    vm1.display()
    vm1.increaseRam()
    vm1.display()
    vm1.increaseDisk(2)
    vm1.display()
}
