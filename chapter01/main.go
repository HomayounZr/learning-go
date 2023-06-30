package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Save the world with GO")

	argsWithProg := os.Args

	num1, err := strconv.Atoi(argsWithProg[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	num2, err := strconv.Atoi(argsWithProg[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	result := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, result)
}
