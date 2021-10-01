package main

import "fmt"

func main() {
	limit := 100
	sumSquares := 0
	squareSum := 0

	for i := 1; i <= limit; i++ {
		sumSquares += i * i
		squareSum += i
	}

	squareSum *= squareSum

	fmt.Println(squareSum - sumSquares)
}
