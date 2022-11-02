/* Alta3 Research | RZFeeser
   cmd.StdoutPipe() - Capturing output from commands   */

package main

import (
    "fmt"
    "io"
    "log"
    "os/exec"
    "strings"
)

func upper(data string) string {

    return strings.ToUpper(data)
}

func main() {
    cmd := exec.Command("ping", "-c", "4", "localhost")

    // this is new -- our command will output through this pipe
    stdout, err := cmd.StdoutPipe()

    if err != nil {
        log.Fatal(err)
    }

    // if we use Start, we must use wait to ensure our command will finish
    if err := cmd.Start(); err != nil {
        log.Fatal(err)
    }

    // grab the stdout from our command and read it in as string data
    data, err := io.ReadAll(stdout)

    if err != nil {
        log.Fatal(err)
    }

    if err := cmd.Wait(); err != nil {
        log.Fatal(err)
    }

    // display the results to the screen
    fmt.Printf("%s\n", upper(string(data)))
}

