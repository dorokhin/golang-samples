package main

import (
    "fmt"
    "strings"
    "time"
)
var c chan struct{} = make(chan struct{})
var StopSpinner bool
var sb strings.Builder

func Spinner(delay time.Duration) {
    for !StopSpinner {
        for _, r := range `-\|/` {

            fmt.Printf( "\rLoading data %s %c", sb.String(), r)
            time.Sleep(delay)
        }
        fmt.Printf( " \r")
    }

    c <- struct{}{}
}

func GetProgress() {
    for i := 0; i < 40; i++ {
        sb.WriteString(".")
        time.Sleep(600 * time.Millisecond)
    }
}

func main() {
    StopSpinner = false
    go Spinner(40 * time.Millisecond)
    go GetProgress()
    time.Sleep(10000 * time.Millisecond)
    StopSpinner = true
    <-c // wait spinner stop

    fmt.Printf(strings.Repeat(" ", 40))
    fmt.Printf( "\rSuccess! 😜")
    time.Sleep(3000 * time.Millisecond)
    fmt.Printf( " \r")
}

