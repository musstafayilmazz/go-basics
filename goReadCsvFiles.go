package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func main() {
	csvfile, err := os.Open("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	df := dataframe.ReadCSV(csvfile)
	fmt.Println(df)

}
