package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	grade, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	days, _ := strconv.Atoi(scanner.Text())

	if days == 0 {
		fmt.Println(20)
	} else if days == 7 {
		fmt.Println(grade)
	} else {
		fmt.Println(math.Max(float64(grade-days), 0))
	}
}
