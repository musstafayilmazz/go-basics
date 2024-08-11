package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content), err)
	fmt.Printf("\n")
	f, err := os.Open("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
