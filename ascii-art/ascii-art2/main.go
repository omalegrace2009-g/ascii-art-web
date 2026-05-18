package main 

import (
	"fmt"
	"strings"
	"os"
)

func Banner(file string) ([]string, error) {
	ban, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	line := strings.Split(string(ban), "\n")
	return line, nil
}

func Get(ban []string, r rune) []string {
	start := (int(r) -32) * 9
	if r < 32 || r > 126 {
		return nil
	}
	return ban[start+1 : start+9]
}

func Print(ban []string, p []string) {
	for _, f := range p {
		if f == "" {
			fmt.Println()
			continue
		}
		for i := 0 ; i <= 7; i++ {
			for _, g := range f {
				char := Get(ban, g)
				if char != nil {
					fmt.Print(char[i])
				}
			}
			fmt.Println()
		}
	}
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Uasge: go run . [string]")
		return
	}
	input := os.Args[1]
	lines := strings.Split(input, "\\n")
	if lines[0] == "" {
		lines = lines[1:]
	}
	g, err := Banner("banner/standard.txt")
	if err != nil {
		fmt.Println("Error reading file: " , err)
	}
	if len(os.Args) > 3 {
		s := os.Args[2] + ".txt"
		h, err := Banner(s)
		if err != nil {
			fmt.Println("Error", err)
			return
		}
		Print(h, lines)
	} else {
		Print(g, lines)
	}
}