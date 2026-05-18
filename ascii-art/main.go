package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [string]")
		return
	}
	input := os.Args[1]
	_, err := ValidateInput(input)
	if err != nil {
		fmt.Println("Errror found", err)
		return
	}
	fit, err := LoadBanner("standard.txt")
	if err != nil {
		fmt.Println("Error reading fille", err)
		return
	}
	 g := GenerateArt(input, fit)
		fmt.Print(g)
	 }
