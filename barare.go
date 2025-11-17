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

	if count%2 == 0 {
		fmt.Println("Bala Barare")
	} else {
		fmt.Println("Payin Barare")
	}
}
