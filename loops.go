package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("i -> %d\t", i)
	}
	fmt.Printf("\n")
	var even = []int{2, 4, 6, 8, 10}

	for index, elem := range even {
		fmt.Printf("%d index = %d\t", index, elem)
	}
}
