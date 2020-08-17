package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range intSlice {
		if value%10 == 0 {
			fmt.Printf("Value %d is even\n", value)
		} else {
			fmt.Printf("Value %d is odd\n", value)
		}
	}
}
