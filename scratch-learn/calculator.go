package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var input1 float64
	var input2 float64
	var Operation string
	fmt.Println()
	fmt.Println("Welcome to Omahlay Calculative World!")
	fmt.Println()
	fmt.Println("Enter first number: ")
	in1, err := fmt.Scan(&input1)
	if err != nil {
		fmt.Println("Input numbers only")
		input11, _ := strconv.Atoi(strings.TrimSuffix(in1, "\n"),in1)
		return
	}
	fmt.Println("Enter second number: ")
	in2, err := fmt.Scan(&input2)
    if err != nil {
		fmt.Println("Input numbers only")
			input22 := strings.TrimSuffix(in2, "\n")
		return
	}
		fmt.Println()
	fmt.Println("Enter Operation: ")
	fmt.Scan(&Operation)
	switch Operation {
	case "+":
		fmt.Println("Result is: ", input11 + input22)
		case "-":
		fmt.Println("Result is: ", input11 - input22)
		case "*":
		fmt.Println("Result is: ", input11 * input22)
		case "/":
		fmt.Println("Result is: ", input11 / input22)
		default:
			fmt.Println("Invalid Input")
	}
}