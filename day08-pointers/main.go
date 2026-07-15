package main

import "fmt"

type User struct { 
    Name string 
    Age int 
}

func increaseByValue(age int) { 
    age = age + 1 
}

func increaseByPointer(age *int) { 
    *age = *age + 1 
}

func birthday(user *User) { 
    user.Age++ 
}

func main() { 
    age := 25

    fmt.Println("Age:", age)
    fmt.Println("Address:", &age)

    agePointer := &age

    fmt.Println("Pointer:", agePointer)
    fmt.Println("Pointer value:", *agePointer)

    increaseByValue(age)
    fmt.Println("After increaseByValue:", age)

    increaseByPointer(&age)
    fmt.Println("After increaseByPointer:", age)

    user := User{
        Name: "Gopher",
        Age:  25,
    }

    birthday(&user)
    fmt.Println("User:", user)

    var emptyPointer *int
    fmt.Println("Is nil:", emptyPointer == nil)

}