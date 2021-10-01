package main

import "fmt"

func sieveOfEratosthenes(limit int) []int {
	result := make([]int, 0, limit)
	prime := make([]bool, limit+1)

	for p := 2; p*p <= limit; p++ {
		if !prime[p] {
			for i := p * p; i <= limit; i += p {
				prime[i] = true
			}
		}
	}

	for i := 2; i <= limit; i++ {
		if !prime[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	limit := 2000000
	sum := 0

	for _, i := range sieveOfEratosthenes(limit) {
		sum += i
	}

	fmt.Println(sum)

}
