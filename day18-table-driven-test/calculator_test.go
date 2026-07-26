package main

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {name: "positive numbers", a: 2, b: 3, expected: 5},
        {name: "negative numbers", a: -2, b: -3, expected: -5},
        {name: "zero", a: 0, b: 5, expected: 5},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Add(test.a, test.b)

            if result != test.expected {
                t.Errorf("expected %d, got %d", test.expected, result)
            }
        })
    }
}

func TestDivide(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
        wantErr  bool
    }{
        {name: "valid division", a: 10, b: 2, expected: 5, wantErr: false},
        {name: "division by zero", a: 10, b: 0, expected: 0, wantErr: true},
        {name: "negative result", a: -10, b: 2, expected: -5, wantErr: false},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Divide(test.a, test.b)

            if test.wantErr {
                if err == nil {
                    t.Fatalf("expected error, got nil")
                }

                return
            }

            if err != nil {
                t.Fatalf("expected nil error, got %v", err)
            }

            if result != test.expected {
                t.Errorf("expected %d, got %d", test.expected, result)
            }
        })
    }
}

func TestGreeting(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected string
    }{
        {name: "with name", input: "Gopher", expected: "Hello, Gopher"},
        {name: "empty name", input: "", expected: "Hello, guest"},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Greeting(test.input)

            if result != test.expected {
                t.Errorf("expected %q, got %q", test.expected, result)
            }
        })
    }
}