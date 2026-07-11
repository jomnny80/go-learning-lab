package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello,", name)
}

func add(a, b int) int {
    return a + b
}

func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b

    return quotient, remainder
}

func printLanguages(languages ...string) {
    for _, language := range languages {
        fmt.Println(language)
    }
}

func main() {
    greet("Gopher")

    sum := add(10, 5)
    fmt.Println("Sum:", sum)

    quotient, remainder := divide(10, 3)
    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)

    printLanguages("Go", "Python", "JavaScript")
}