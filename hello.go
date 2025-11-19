package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count := scanner.Text()

	fmt.Println("Hello CodeCup " + count + "!")
}
