package main

import (
    "fmt"
    "errors"
)

func Hello(name string) (string, error) {
    if name == "" {
        return name, errors.New("empty name")
    }

    message := fmt.Sprintf("Hello there %v", name)
    return message, nil
}
