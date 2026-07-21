package main

import (
    "fmt"
    "sync"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func printNumber(number int, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Println("Number:", number)
}

func main() {
    go sayHello()

    fmt.Println("Hello from main")

    time.Sleep(1 * time.Second)

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1)

        go printNumber(i, &wg)
    }

    wg.Wait()

    fmt.Println("All goroutines finished")
}