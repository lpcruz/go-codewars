package main

import (
	"fmt"
	"os"
	"strings"
)

func Feast(beast string, dish string) bool {
	normalize := func(s string) string {
		return strings.ReplaceAll(strings.TrimSpace(s), " ", "")
	}

	getFirstAndLast := func(s string) (first, last string) {
		list := strings.Split(s, "")
		return list[0], list[len(list)-1]
	}

	// Normalize and split the strings
	normalizedBeast := normalize(beast)
	normalizedDish := normalize(dish)

	// Get first and last characters of each list
	beastFirstChar, beastLastChar := getFirstAndLast(normalizedBeast)
	dishFirstChar, dishLastChar := getFirstAndLast(normalizedDish)

	// Compare for deep equality
	return beastFirstChar == dishFirstChar && beastLastChar == dishLastChar
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go 'beast' 'dish'")
		os.Exit(1)
	}
	beast := os.Args[1]
	dish := os.Args[2]
	result := Feast(beast, dish)
	fmt.Println(result)
}
