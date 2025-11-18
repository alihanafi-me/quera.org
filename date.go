package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	fmt.Println("saal:" + string(input[0]) + string(input[1]))
	fmt.Println("maah:" + string(input[2]) + string(input[3]))
}
