package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(number int) bool {
	if number == 1 {
		return false
	} else if number == 2 {
		return true
	}

	sqrt := int(math.Sqrt(float64(number)))

	for i := 2; i <= sqrt; i++ {
		if number%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	end, _ := strconv.Atoi(scanner.Text())

	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
