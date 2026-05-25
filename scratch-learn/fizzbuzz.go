package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var g int
	fmt.Println("Enter a number:")
	read := bufio.NewReader(os.Stdin)
		_, err := fmt.Scan(&g)
		if err != nil {
			fmt.Println("input numbers only!!")
		_, _ = read.ReadString('\n')
			return 
		}
		for i := 1; i <= g; i++ {
	if i%5 == 0 && i%3 == 0 {
		fmt.Println("fizzbuzz")
	} else if i%3 == 0 {
		fmt.Println("fizz")
	} else if i%5 == 0 {
		fmt.Println("buzz")
	}  else {
		fmt.Println(i)
	}
	}
}