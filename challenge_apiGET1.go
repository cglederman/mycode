package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"
)

type SpaceX []struct {
        CapsuleSerial      string    `json:"capsule_serial"`
        CapsuleID          string    `json:"capsule_id"`
        Status             string    `json:"status"`
        OriginalLaunch     time.Time `json:"original_launch"`
        OriginalLaunchUnix int       `json:"original_launch_unix"`
        Missions           []struct {
                Name   string `json:"name"`
                Flight int    `json:"flight"`
        } `json:"missions"`
        Landings   int    `json:"landings"`
        Type       string `json:"type"`
        Details    string `json:"details"`
        ReuseCount int    `json:"reuse_count"`
}

func main() {
        // define our URL (API) as a string
        url := "https://api.spacexdata.com/v3/capsules"

        // Build the request
        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
                log.Fatal("NewRequest: ", err)
                return
        }
        client := &http.Client{}

         resp, err := client.Do(req)
         if err != nil {
                log.Fatal("Do: ", err)
                return
        }

        defer resp.Body.Close()

        var record SpaceX

        if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
                log.Println(err)
        }

        for launchNo, launchData := range record {
            fmt.Println("Capsule Record     =\n", launchData)
            fmt.Println("Record Number      =", launchNo)
            fmt.Println("Capsule Serial     =", launchData.CapsuleSerial)
            fmt.Println("Capsule ID         =", launchData.CapsuleID)
            fmt.Println("Original Launch    =", launchData.OriginalLaunch)
            fmt.Println("Landings           =", launchData.Landings)
            fmt.Println("Type               =", launchData.Type)
            fmt.Println("Details            =", launchData.Details)
            fmt.Println("Reuse Count        =", launchData.ReuseCount)
            fmt.Println("End of Record\n\n")

        }
}


