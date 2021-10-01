package main

import "fmt"

func reverse(n int) int {
	rev := 0

	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}

	return rev
}

func main() {
	largestPalindromic := 0
	limit := 1000

	for x := 100; x < limit; x++ {
		for y := x; y < limit; y++ {
			tmpResult := x * y

			if tmpResult < largestPalindromic {
				continue
			}

			if tmpResult == reverse(tmpResult) {
				largestPalindromic = tmpResult
			}

		}
	}

	fmt.Println(largestPalindromic)
}
