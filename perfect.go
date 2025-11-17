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
	count, _ := strconv.Atoi(scanner.Text())
	sum := 1

	for i := 2; i <= int(math.Sqrt(float64(count))); i++ {
		if count%i == 0 {
			sum += i

			if count/i != i {
				sum += count / i
			}
		}
	}

	if sum == count {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
