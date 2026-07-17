package main

import (
    "fmt"

    "go-learning-lab/day10-packages/calculator"
)

func main() {
    sum := calculator.Add(10, 5)
    difference := calculator.Subtract(10, 5)

    fmt.Println("Sum:", sum)
    fmt.Println("Difference:", difference)
}