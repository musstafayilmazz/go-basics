package main

import (
	"fmt"
	"log"
	"rsc.io/pdf"
)

func main() {

	file, err := pdf.Open("asd.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file.NumPage())
	fmt.Println(file.Page(1).Content())
}
