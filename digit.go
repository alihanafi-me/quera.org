package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	result := count % 9

	if result == 0 {
		fmt.Println("9")
	} else {
		fmt.Println(result)
	}
}
