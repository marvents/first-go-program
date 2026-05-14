package Ops

import "fmt"
import "program/Logs"

func Add(a, b float32) float32 {
	return a + b
}

func Subtract(a, b float32) float32 {
	return a - b
}

func Multiply(a, b float32) float32 {
	return a * b
}

func Divide(a, b float32) float32 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		return 0
	}
	return a / b
}

func CalcGui(a, b, result *float32, op *string, logs *[]string) {
	fmt.Println("==== Calculator ====")
	fmt.Print("Enter first number: ")
	fmt.Scan(a)
	fmt.Print("Enter second number: ")
	fmt.Scan(b)
	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(op)

	switch *op {
	case "+":
		*result = Add(*a, *b)
	case "-":
		*result = Subtract(*a, *b)
	case "*":
		*result = Multiply(*a, *b)
	case "/":
		*result = Divide(*a, *b)
	default:
		fmt.Println("Invalid operation")
		return
	}
	Logs.LogOperation(*a, *b, *op, *result, logs)
}
