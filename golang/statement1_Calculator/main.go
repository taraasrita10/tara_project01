package main

import (
	"fmt"
)

func calculator(op1, op2 float64, op string) float64 {
	switch op {
	case "+":
		return op1 + op2
	case "-":
		return op1 - op2
	case "*":
		return op1 * op2
	case "/":
		if op2 == 0 {
			return 0
		}
		return op1 / op2
	default:
		return 0
	}
}

func main() {
	var op1, op2 float64
	var op string
	var use string
	fmt.Println("Would you like to use the calculator? (y/n)")
	fmt.Scan(&use)
	for use != "n" {
		fmt.Println("Enter the first operand : ")
		fmt.Scan(&op1)
		fmt.Println("Enter the second operand : ")
		fmt.Scan(&op2)
		fmt.Println("Enter the operator : ")
		fmt.Scan(&op)
		result := calculator(op1, op2, op)
		if result == 0 {
			if op2 == 0 && op == "/" {
				fmt.Println("Division with 0 is not possible.")
			} else if op != "+" && op != "-" && op != "*" && op != "/" {
				fmt.Println("Operator not valid")
			} else {
				fmt.Println(result)
			}
		} else {
			fmt.Println(result)
		}
		fmt.Println("Would you like to continue using the calculator? (y/n)")
		fmt.Scan(&use)
		if use != "y" && use != "n" {
			break
		}
	}
}
