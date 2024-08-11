package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	fmt.Println(slice)

	m := make([]byte, 2, 3)
	fmt.Println(m)
}
