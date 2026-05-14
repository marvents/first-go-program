package Ops

import "fmt"

func Add(a, b float32) float32 { // هدا جزء لم ادرسه بعد لاكن فهمته تمت كتابته من قبل auto completion وما فهمت هو (a,b float32) تعني type خاص بالمتغييرين اما float32 ثانية تعني type خاص ب return وهدا تعلمته من قبل في typescript
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