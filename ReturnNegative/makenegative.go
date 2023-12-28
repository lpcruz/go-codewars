package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func MakeNegative(x int) int {
	isNegative := math.Signbit(float64(x))
	if isNegative {
		return x
	}
	return -x
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run makenegative.go <number>")
		return
	}
	arg := os.Args[1]
	x, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid number:", arg)
		return
	}
	result := MakeNegative(x)
	fmt.Println(result)
}
