package main

import "fmt"

type Speaker interface {
    Speak() string
}

type User struct {
    Name string
}

func (u User) Speak() string {
    return "Hello, I am " + u.Name
}

type Robot struct {
    ID string
}

func (r Robot) Speak() string {
    return "Robot " + r.ID + " online"
}

func printMessage(s Speaker) {
    fmt.Println(s.Speak())
}

func printAnything(value any) {
    fmt.Println(value)
}

func main() {
    user := User{Name: "Gopher"}
    robot := Robot{ID: "R2"}

    printMessage(user)
    printMessage(robot)

    printAnything("Go")
    printAnything(100)
    printAnything(true)

    var value any = "Learning interface"

    text, ok := value.(string)
    if ok {
        fmt.Println("Text:", text)
    }
}