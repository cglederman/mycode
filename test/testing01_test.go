package main

import (
    "testing"
    "regexp"
    "fmt"
    "strings"
)

func TestHello(t *testing.T) {
    testName := "RZFeeser"
    desiredResult := regexp.MustCompile("Hello there "+testName)

    result, err := Hello("RZFeeser")
    if !desiredResult.MatchString(result) || err != nil {
        t.Fatalf("Wanted %v, but got, %v, or got %v", desiredResult, result, err)
    }
}
func TestHelloNoName(t *testing.T) {
    result, err := Hello("")
    fmt.Println(err)
    fmt.Println(err.Error())
    if  result != "" || err == nil || !strings.Contains(err.Error(), "empty name") {
        t.Fatalf("Test failed to produce correct failure for empty string. Got, %v", result)
    }
}
