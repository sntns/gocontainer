package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Worker processed: (no input)\n")
		return
	}
	
	input := strings.Join(os.Args[1:], " ")
	fmt.Printf("Worker processed: %s\n", strings.ToUpper(input))
}