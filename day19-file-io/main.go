package main

import (
    "bufio"
    "fmt"
    "os"
)

const fileName = "day19-file-io/notes.txt"

func writeFile() error {
    content := []byte("Learning Go File I/O\nReading and writing files")

    return os.WriteFile(fileName, content, 0644)
}

func readFile() error {
    content, err := os.ReadFile(fileName)
    if err != nil {
        return err
    }

    fmt.Println("File content:")
    fmt.Println(string(content))

    return nil
}

func appendFile() error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString("\nNew line appended by Go")
    return err
}

func readLineByLine() error {
    file, err := os.Open(fileName)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    fmt.Println("Read line by line:")

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    return scanner.Err()
}

func main() {
    if err := writeFile(); err != nil {
        fmt.Println("Write error:", err)
        return
    }

    if err := appendFile(); err != nil {
        fmt.Println("Append error:", err)
        return
    }

    if err := readFile(); err != nil {
        fmt.Println("Read error:", err)
        return
    }

    if err := readLineByLine(); err != nil {
        fmt.Println("Read line error:", err)
        return
    }
}
