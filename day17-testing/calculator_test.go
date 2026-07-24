package main

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(10, 2)

    if err != nil {
        t.Fatalf("expected nil error, got %v", err)
    }

    expected := 5

    if result != expected {
        t.Errorf("expected %d, got %d", expected, result)
    }
}

func TestDivideByZero(t *testing.T) {
    _, err := Divide(10, 0)

    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}