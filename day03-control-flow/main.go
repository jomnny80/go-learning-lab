package main

import "fmt"

func main() {
    score := 82

    if score >= 60 {
        fmt.Println("Passed")
    } else {
        fmt.Println("Failed")
    }

    switch {
    case score >= 90:
        fmt.Println("Grade A")
    case score >= 80:
        fmt.Println("Grade B")
    case score >= 70:
        fmt.Println("Grade C")
    case score >= 60:
        fmt.Println("Grade D")
    default:
        fmt.Println("Failed")
    }

    for i := 1; i <= 5; i++ {
        fmt.Println("Count:", i)
    }

    languages := []string{"Go", "Python", "JavaScript"}

    for index, language := range languages {
        fmt.Println(index, language)
    }
}