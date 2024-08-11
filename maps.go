package main

import "fmt"

func main() {
	fmt.Println("### Maps ###")

	mydict := make(map[string]int)
	general_dict := make(map[int]map[string]int)

	mydict["age"] = 24
	mydict["location"] = 3
	general_dict[0] = mydict
	fmt.Println(mydict)
	fmt.Println(general_dict)
	fmt.Println(general_dict[0]["age"])

}
