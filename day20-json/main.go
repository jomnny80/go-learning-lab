package main

import (
    "encoding/json"
    "fmt"
    "os"
)

const fileName = "day20-json/users.json"

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}

func encodeUser() error {
    user := User{
        ID:    1,
        Name:  "Gopher",
        Email: "gopher@example.com",
    }

    data, err := json.MarshalIndent(user, "", "  ")
    if err != nil {
        return err
    }

    fmt.Println("Single user JSON:")
    fmt.Println(string(data))

    return nil
}

func decodeUser() error {
    jsonData := []byte(`{
        "id": 2,
        "name": "Go Learner",
        "email": "learner@example.com"
    }`)

    var user User

    err := json.Unmarshal(jsonData, &user)
    if err != nil {
        return err
    }

    fmt.Println("Decoded user:")
    fmt.Println(user)

    return nil
}

func writeUsersToFile() error {
    users := []User{
        {
            ID:    1,
            Name:  "Gopher",
            Email: "gopher@example.com",
        },
        {
            ID:   2,
            Name: "Go Learner",
        },
    }

    data, err := json.MarshalIndent(users, "", "  ")
    if err != nil {
        return err
    }

    return os.WriteFile(fileName, data, 0644)
}

func readUsersFromFile() error {
    data, err := os.ReadFile(fileName)
    if err != nil {
        return err
    }

    var users []User

    err = json.Unmarshal(data, &users)
    if err != nil {
        return err
    }

    fmt.Println("Users from file:")

    for _, user := range users {
        fmt.Printf("ID: %d, Name: %s, Email: %s\n", user.ID, user.Name, user.Email)
    }

    return nil
}

func main() {
    if err := encodeUser(); err != nil {
        fmt.Println("Encode user error:", err)
        return
    }

    if err := decodeUser(); err != nil {
        fmt.Println("Decode user error:", err)
        return
    }

    if err := writeUsersToFile(); err != nil {
        fmt.Println("Write users error:", err)
        return
    }

    if err := readUsersFromFile(); err != nil {
        fmt.Println("Read users error:", err)
        return
    }
}
