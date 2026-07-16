package main

import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }

    return a / b, nil
}

func findUser(id int) (string, error) {
    if id != 1 {
        return "", fmt.Errorf("user with id %d not found", id)
    }

    return "Gopher", nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)

    user, err := findUser(2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("User:", user)
}
