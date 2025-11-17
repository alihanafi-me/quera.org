package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	maxNumber := 0
	num := 0

	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)

	for i := 0; i < count; i++ {
		num, _ = strconv.Atoi(fields[i])
		maxNumber = int(math.Max(float64(maxNumber), float64(num)))
	}

	fmt.Println(maxNumber)
}
