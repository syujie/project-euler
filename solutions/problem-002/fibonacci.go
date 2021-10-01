package main

import "fmt"

func main() {
	x := 1
	y := 2
	t := 0

	for x <= 4000000 {
		if x%2 == 0 {
			t += x
		}

		z := x + y
		x = y
		y = z
	}

	fmt.Println(t)
}
