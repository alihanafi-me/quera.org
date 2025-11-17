package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	number := new(big.Int)
	number.SetString(input, 10)
	log2 := big.NewInt(int64(number.BitLen()))
	base := big.NewInt(2)

	fmt.Println(new(big.Int).Exp(base, log2, nil))
}
