package main

import "fmt"

func main() {
    numbers := [3]int{10, 20, 30}

    fmt.Println("Array:", numbers)
    fmt.Println("First number:", numbers[0])

    scores := []int{80, 90, 100}
    scores = append(scores, 95)

    fmt.Println("Slice:", scores)
    fmt.Println("Length:", len(scores))
    fmt.Println("Capacity:", cap(scores))

    languages := []string{"Go", "Python", "JavaScript"}

    for index, language := range languages {
        fmt.Println(index, language)
    }

    profile := map[string]string{
        "name":     "Gopher",
        "language": "Go",
    }

    profile["level"] = "Beginner"
    profile["name"] = "Go Learner"

    fmt.Println("Profile:", profile)

    level, ok := profile["level"]
    if ok {
        fmt.Println("Level:", level)
    }

    delete(profile, "level")

    for key, value := range profile {
        fmt.Println(key, value)
    }
}