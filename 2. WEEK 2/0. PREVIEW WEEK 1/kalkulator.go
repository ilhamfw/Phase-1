package main

import (
    "fmt"
    "os"
    "strconv"
)

func kalkulator() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
        return
    }

    operation := os.Args[1]
    operand1, err1 := strconv.ParseFloat(os.Args[2], 64)
    operand2, err2 := strconv.ParseFloat(os.Args[3], 64)

    if err1 != nil || err2 != nil {
        fmt.Println("Invalid operands. Please enter valid numbers.")
        return
    }

    var result float64

    switch operation {
    case "add":
        result = operand1 + operand2
    case "subtract":
        result = operand1 - operand2
    case "multiply":
        result = operand1 * operand2
    case "divide":
        if operand2 == 0 {
            fmt.Println("Cannot divide by zero.")
            return
        }
        result = operand1 / operand2
    default:
        fmt.Println("Invalid operation. Supported operations: add, subtract, multiply, divide.")
        return
    }

    fmt.Printf("Result: %.2f\n", result)
}
