package main

import (
	"fmt"
)

func main() {
	var g int
	fmt.Println("Enter a number:")
		fmt.Scan(&g)
		for i := 1; i <= g; i++ {
	if i%5 == 0 && i%3 == 0 {
		fmt.Println("fizzbuzz")
		continue
	}
	if i%3 == 0 {
		fmt.Println("fizz")
		continue
	}
	if i%5 == 0 {
		fmt.Println("buzz")
continue
	}  else {
		fmt.Println(i)
	}
	}
}