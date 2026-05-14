package Logs

import "fmt"

func LogOperation(a, b float32, op string, result float32, logs *[]string) {
	*logs = append(*logs, fmt.Sprintf("%.2f %s %.2f = %.2f", a, op, b, result))
}