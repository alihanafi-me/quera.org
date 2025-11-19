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
	first, _ := strconv.ParseFloat(scanner.Text(), 64)

	scanner.Scan()
	second, _ := strconv.ParseFloat(scanner.Text(), 64)

	bmi := math.Round((100*first)/(second*second)) / 100
	fmt.Printf("%.2f", bmi)
	fmt.Println()

	if bmi < 18.5 {
		fmt.Println("Underweight")
	} else if bmi < 25 {
		fmt.Println("Normal")
	} else if bmi < 30 {
		fmt.Println("Overweight")
	} else {
		fmt.Println("Obese")
	}
}
