package main

import (
    "fmt"
    "time"
)

func sendMessage(ch chan string, message string, delay time.Duration) {
    time.Sleep(delay)

    ch <- message
}

func slowTask(ch chan string) {
    time.Sleep(3 * time.Second)

    ch <- "task finished"
}

func sendNumbers(ch chan int) {
    for i := 1; i <= 3; i++ {
        ch <- i
        time.Sleep(500 * time.Millisecond)
    }

    close(ch)
}

func worker(done chan bool) {
    for {
        select {
        case <-done:
            fmt.Println("worker stopped")
            return
        default:
            fmt.Println("working...")
            time.Sleep(500 * time.Millisecond)
        }
    }
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go sendMessage(ch1, "message from ch1", 2*time.Second)
    go sendMessage(ch2, "message from ch2", 1*time.Second)

    select {
    case message := <-ch1:
        fmt.Println(message)
    case message := <-ch2:
        fmt.Println(message)
    }

    timeoutCh := make(chan string)

    go slowTask(timeoutCh)

    select {
    case message := <-timeoutCh:
        fmt.Println(message)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout")
    }

    emptyCh := make(chan string)

    select {
    case message := <-emptyCh:
        fmt.Println(message)
    default:
        fmt.Println("no message received")
    }

    numbers := make(chan int)

    go sendNumbers(numbers)

    for {
        number, ok := <-numbers
        if !ok {
            fmt.Println("numbers channel closed")
            break
        }

        fmt.Println("Number:", number)
    }

    done := make(chan bool)

    go worker(done)

    time.Sleep(2 * time.Second)

    done <- true

    time.Sleep(1 * time.Second)
}