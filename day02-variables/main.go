package main

import "fmt"

func main() {
    var name string = "Gopher"
    var age int = 25
    price := 99.9
    isLearning := true

    const language = "Go"

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
    fmt.Println("Price:", price)
    fmt.Println("Learning:", isLearning)
    fmt.Println("Language:", language)

    var emptyString string
    var zeroInt int
    var zeroFloat float64
    var defaultBool bool

    fmt.Println("Empty string:", emptyString)
    fmt.Println("Zero int:", zeroInt)
    fmt.Println("Zero float:", zeroFloat)
    fmt.Println("Default bool:", defaultBool)
}