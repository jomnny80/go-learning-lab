package main

import "fmt"

func simpleDefer() {
    defer fmt.Println("simpleDefer end")

    fmt.Println("simpleDefer start")
}

func multipleDefer() {
    defer fmt.Println("First defer")
    defer fmt.Println("Second defer")
    defer fmt.Println("Third defer")

    fmt.Println("multipleDefer start")
}

func earlyReturn(success bool) {
    defer fmt.Println("cleanup")

    if !success {
        fmt.Println("failed")
        return
    }

    fmt.Println("succeeded")
}

func argumentEvaluation() {
    name := "Go"

    defer fmt.Println("Deferred name:", name)

    name = "Gopher"

    fmt.Println("Current name:", name)
}

func deferredFunction() {
    name := "Go"

    defer func() {
        fmt.Println("Deferred function name:", name)
    }()

    name = "Gopher"

    fmt.Println("Current name:", name)
}

func main() {
    simpleDefer()

    fmt.Println("---")

    multipleDefer()

    fmt.Println("---")

    earlyReturn(false)

    fmt.Println("---")

    argumentEvaluation()

    fmt.Println("---")

    deferredFunction()
}