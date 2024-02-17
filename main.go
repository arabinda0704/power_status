package main

import (
    "fmt"
    "time"
    "github.com/goprojs/power_status/utils"
)

func main() {
    initialState, err := utils.GetAndSend("config.json")
    // Send initial status to server
    if err != nil {
        fmt.Println("Error sending initial data:", err)
    }
    // Print initial battery status
    fmt.Println("Initial battery status:", initialState)
    for {
        currentState, err := utils.BatteryHasPowerSupply()
        if err != nil {
            fmt.Println(err)
        }
        if currentState != initialState {
            initialState = currentState
            _, err := utils.GetAndSend("config.json")
            if err != nil {
                fmt.Println("Error sending initial data:", err)
            }
            time.Sleep(1 * time.Second)
        }
    }
}
