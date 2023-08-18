package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var u int = 0

func Sqrt(x float64) float64 {
	z := 1.0
	previousZ := 0.0
	epsilon := 1e-10

	for i := 0; i < 100; i++ {
		u++
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-previousZ) < epsilon {
			break
		}
		previousZ = z
	}
	return z
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}

	x, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("There was an error converting the string!", err)
		return
	}

	fmt.Println("Custom Sqrt:", Sqrt(x), "Times looped until convergence:", u)
	fmt.Println("math.Sqrt:", math.Sqrt(x))
	time.Sleep(5 * time.Second)
}
