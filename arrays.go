package main

import "fmt"

func main() {
	fmt.Println("### Arrays ###")

	var arr [2]int

	arr[0] = 5
	arr[1] = 10
	fmt.Println(arr, len(arr), arr[0])

	var even = [4]int{2, 4, 6, 8}
	fmt.Printf("even's last member is %d\n", even[len(even)-1])
}
