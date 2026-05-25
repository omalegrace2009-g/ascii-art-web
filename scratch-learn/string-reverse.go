package main
import (
	"bufio"
	"os"
	"strings"
	"fmt"
)
func main() {
	fmt.Println("Enter a word")
read := bufio.NewReader(os.Stdin)
in, _ := read.ReadString('\n')
g := strings.TrimSuffix(in, "\n")
	fmt.Print("Reversed string: ")
for i := len(g)-1; i >= 0; i-- {
	fmt.Print(string(in[i]))
}
fmt.Println()
}