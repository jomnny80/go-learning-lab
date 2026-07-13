package main

import "fmt"

type Status string

const (
    StatusActive   Status = "active"
    StatusInactive Status = "inactive"
)

type Address struct {
    City    string
    Country string
}

type User struct {
    Name    string
    Age     int
    Email   string
    Status  Status
    Address Address
}

func printUser(user User) {
    fmt.Println("Name:", user.Name)
    fmt.Println("Age:", user.Age)
    fmt.Println("Email:", user.Email)
    fmt.Println("Status:", user.Status)
    fmt.Println("City:", user.Address.City)
    fmt.Println("Country:", user.Address.Country)
}

func main() {
    user := User{
        Name:   "Gopher",
        Age:    25,
        Email:  "gopher@example.com",
        Status: StatusActive,
        Address: Address{
            City:    "Taipei",
            Country: "Taiwan",
        },
    }

    printUser(user)

    user.Age = 26
    user.Status = StatusInactive

    fmt.Println("Updated user:")
    printUser(user)

    var emptyUser User
    fmt.Println("Empty user:", emptyUser)
}