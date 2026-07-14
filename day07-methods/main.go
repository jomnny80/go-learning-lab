package main

import "fmt"

type User struct {
    Name  string
    Age   int
    Email string
}

func (u User) printProfile() {
    fmt.Println("Name:", u.Name)
    fmt.Println("Age:", u.Age)
    fmt.Println("Email:", u.Email)
}

func (u User) isAdult() bool {
    return u.Age >= 18
}

func (u *User) birthday() {
    u.Age++
}

func (u *User) updateEmail(email string) {
    u.Email = email
}

type Score int

func (s Score) isPassed() bool {
    return s >= 60
}

func main() {
    user := User{
        Name:  "Gopher",
        Age:   25,
        Email: "gopher@example.com",
    }

    user.printProfile()

    fmt.Println("Is adult:", user.isAdult())

    user.birthday()
    user.updateEmail("new-gopher@example.com")

    fmt.Println("Updated user:")
    user.printProfile()

    score := Score(85)
    fmt.Println("Passed:", score.isPassed())
}