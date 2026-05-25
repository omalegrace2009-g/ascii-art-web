package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"	
)

func main() {
	fmt.Println("Enter a sentence")
	read := bufio.NewReader(os.Stdin)
	n, _ := read.ReadString('\n')
	fit := strings.TrimSuffix(n, "\n")
	fiit := strings.ToLower(fit)
	lit := strings.ReplaceAll(fiit, " ", "")
	reversed := ""
	for i := len(lit)-1; i >= 0; i-- {
		reversed += string(lit[i])
	}
	if reversed != lit {
		fmt.Println("Input is not a palindrome")
	} else {
		fmt.Println("Input is a palindrome")
	}
}