package main

import (
    "fmt"
    "math/rand"
    "net/http"
    _ "net/http/pprof"
    "time"
)

func doSomething() {
    for i := 0; i < 100; i++ {
        rand.Intn(100)
    }
}

func main() {
    go func() {
        fmt.Println(http.ListenAndServe("localhost:8080", nil))

    }()
    for i := 0; i < 100; i++ {
        go doSomething()
    }

    time.Sleep(200 * time.Second)
}