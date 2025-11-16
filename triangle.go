package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numbers := []int{0, 0, 0}

	numbers[0], _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	numbers[1], _ = strconv.Atoi(scanner.Text())
	scanner.Scan()
	numbers[2], _ = strconv.Atoi(scanner.Text())

	sort.Ints(numbers)

	if math.Pow(float64(numbers[0]), 2)+math.Pow(float64(numbers[1]), 2) == math.Pow(float64(numbers[2]), 2) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
