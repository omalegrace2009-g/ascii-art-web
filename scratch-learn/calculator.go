package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var input1 float64
	var input2 float64
	var Operation string
	fmt.Println()
	fmt.Println("Welcome to Omahlay Calculative World!")
	fmt.Println()
	fmt.Println("Enter first number: ")
	read := bufio.NewReader(os.Stdin)
	_, err := fmt.Scan(&input1)
     if err != nil {
     fmt.Println("Input numbers only!!")
	 _, _ = read.ReadString('\n')
     return
     }
	fmt.Println("Enter second number: ")
	_, err = fmt.Scan(&input2)
    if err != nil {
		fmt.Println("Input numbers only!!")
		_, _ = read.ReadString('\n')
		return
	}
	fmt.Println("Enter Operation: ")
	fmt.Scan(&Operation)
	switch Operation {
	case "+":
		fmt.Println("Result is: ", input1 + input2)
		case "-":
		fmt.Println("Result is: ", input1 - input2)
		case "*":
		fmt.Println("Result is: ", input1 * input2)
		case "/":
			if input2 == 0 {
			fmt.Println("cannot divide by 0")
			return
		}
		fmt.Println("Result is: ", input1 / input2)
		default:
			fmt.Println("Invalid Input")
	}
}