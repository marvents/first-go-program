package main

import "fmt"
import "program/Ops"

func main() {
	fmt.Println("==== Calculator ====")
	var a, b float32
	var op string
	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Print("Enter operation (+, -, *, /): ")
	fmt.Scan(&op)
	var result float32

	switch op {
		case "+":
			result = Ops.Add(a, b)
		case "-":
			result = Ops.Subtract(a, b)
		case "*":
			result = Ops.Multiply(a, b)
		case "/":
			result = Ops.Divide(a, b)
		default:
			fmt.Println("Invalid operation")
			return

	}
	fmt.Printf("Result: %.2f\n", result) // هنا ايضا تمت اضافة هدا الجزء من قبل auto completion %.2f\n وهدا يعني انه يريد رقمين بعد فاصلة من نوع float و endl; هه كما في cpp
}