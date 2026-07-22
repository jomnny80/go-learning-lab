package main

import "fmt"

func worker(ch chan string) {
    ch <- "work done"
}

func calculate(ch chan int) {
    result := 10 + 5

    ch <- result
}

func sendNumbers(ch chan int) {
    for i := 1; i <= 3; i++ {
        ch <- i
    }

    close(ch)
}

func sendMessage(ch chan<- string) {
    ch <- "hello from send-only channel"
}

func receiveMessage(ch <-chan string) {
    message := <-ch

    fmt.Println(message)
}

func main() {
    messageCh := make(chan string)

    go worker(messageCh)

    message := <-messageCh
    fmt.Println(message)

    resultCh := make(chan int)

    go calculate(resultCh)

    result := <-resultCh
    fmt.Println("Result:", result)

    bufferedCh := make(chan string, 2)

    bufferedCh <- "first"
    bufferedCh <- "second"

    fmt.Println(<-bufferedCh)
    fmt.Println(<-bufferedCh)

    numbers := make(chan int)

    go sendNumbers(numbers)

    for number := range numbers {
        fmt.Println("Number:", number)
    }

    directionCh := make(chan string)

    go sendMessage(directionCh)

    receiveMessage(directionCh)
}