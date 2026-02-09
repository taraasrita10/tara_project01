package main

import (
	"fmt"
	"strings"
)

//Calculator function :
// Takes two operators and an operand as input, performs
// a mathematical function and provides the output.
// Involves addition, subtraction, multiplication and division of float64 numbers.

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
	strings.ToLower(use)
	for use == "y" {
		fmt.Println("Enter the first operand : ")
		_, err1 := fmt.Scan(&op1)
		fmt.Println("Enter the second operand : ")
		_, err2 := fmt.Scan(&op2)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid Operand!")
			continue
		}
		fmt.Println("Enter the operator : ")
		fmt.Scan(&op)
		if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println(("Invalid Operator!"))
			continue
		}
		if op2 == 0 && op == "/" {
			fmt.Println("Division by 0 is not possible")
			continue
		}
		result := calculator(op1, op2, op)
		fmt.Printf("%g %s %g = %g\n", op1, op, op2, result)
		fmt.Println("Would you like to continue using the calculator? (y/n)")
		fmt.Scan(&use)
		strings.ToLower(use)
		if use != "y" && use != "n" {
			break
		}
	}
}
