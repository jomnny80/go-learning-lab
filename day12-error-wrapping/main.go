package main

import (
    "errors"
    "fmt"
)

var ErrUserNotFound = errors.New("user not found")

type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func findUser(id int) error {
    if id != 1 {
        return fmt.Errorf("find user by id %d: %w", id, ErrUserNotFound)
    }

    return nil
}

func validateName(name string) error {
    if name == "" {
        return fmt.Errorf("validate user input: %w", &ValidationError{
            Field:   "name",
            Message: "cannot be empty",
        })
    }

    return nil
}

func main() {
    err := findUser(2)

    if errors.Is(err, ErrUserNotFound) {
        fmt.Println("User does not exist")
    }

    err = validateName("")

    var validationErr *ValidationError

    if errors.As(err, &validationErr) {
        fmt.Println("Validation error")
        fmt.Println("Field:", validationErr.Field)
        fmt.Println("Message:", validationErr.Message)
    }
}