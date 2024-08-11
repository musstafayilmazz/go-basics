package main

import "fmt"

var job string = "Data Scientist"

func main() {
	var foo int
	foo = 6
	boo := 13
	var coo int = 14
	fmt.Printf("%T %T %T\n", foo, boo, coo)
	fmt.Println(foo + boo + coo)

	var name string = "Mustafa"
	fmt.Println(name + " " + job)
	fmt.Printf("My name is %v and my job is %v \n", name, job)
}
