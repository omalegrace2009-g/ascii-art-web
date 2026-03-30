package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println()
		fmt.Println("Welcome To OMAHLAY calculative World!")
		time.Sleep(3 * time.Second)
		var operation string
		fmt.Println("Choose operation: ")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		fmt.Println("5. Quit")
		fmt.Println("6. Help")
		fmt.Scanln(&operation)
		var number1 int
		var number2 int
		switch operation {
		case "1":
		start:
			fmt.Println("Enter firstnumber ")
			_, err := fmt.Scan(&number1)
			fmt.Println("Enter second number: ")
			_, err1 := fmt.Scan(&number2)
			if err != nil || err1 != nil {
				fmt.Println("input valid numbers only! ")
				goto start
			}
			fmt.Println(number1 + number2)

		case "2":

			fmt.Println("Enter firstnumber ")
			_, err := fmt.Scan(&number1)
			fmt.Println("Enter second number: ")
			_, err1 := fmt.Scan(&number2)
			if err != nil || err1 != nil {
				fmt.Println("input valid numbers only! ")
			}
			fmt.Println(number1 - number2)

		case "3":

			fmt.Println("Enter firstnumber ")
			_, err := fmt.Scan(&number1)
			fmt.Println("Enter second number: ")
			_, err1 := fmt.Scan(&number2)
			if err != nil || err1 != nil {
				fmt.Println("input valid numbers only! ")
			}
			fmt.Println(number1 * number2)

		case "4":

			fmt.Println("Enter firstnumber ")
			_, err := fmt.Scan(&number1)
			fmt.Println("Enter second number: ")
			_, err1 := fmt.Scan(&number2)
			if err != nil || err1 != nil {
				fmt.Println("input valid numbers only! ")
			}
			fmt.Println(number1 / number2)
			if number2 == 0 {
				fmt.Println("division by Zero is undefined")
			}

		case "5":
			fmt.Println("program quit successfully..")
			break
		case "6":
			fmt.Println("input the numbers you want to work to carry out the operation on")
			fmt.Println("choose the operation you want e.g select <<1>> to add ")
			fmt.Println("enter option <<5>> to quit the program and <<6>> to see how to operate")
		default:
			fmt.Println("Choose from the options listed")
		}
	}
}
