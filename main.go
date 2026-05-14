package main

import "fmt"
import "program/Ops"

func main() {
	var (
	logs []string
	a, b, result float32
	op string
	forOps int
	)

	for {
		fmt.Println("1. Calculator")
		fmt.Println("2. View Logs")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		fmt.Scan(&forOps)

		switch forOps {
			case 1:
				Ops.CalcGui(&a, &b, &result, &op, &logs)
				fmt.Printf("result: %.2f\n", result)
			case 2:
				fmt.Println(logs)
			case 3:
				return
			default:
				fmt.Println("Invalid operation")
		}
		fmt.Println("=====================")
	}



}