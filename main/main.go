package main

import (
	"fmt"

	"csvHelper/fixCsvSize"
	"csvHelper/loadCsv"
)

func main() {
	file := "../example.csv"

	// First, loads the CSV data, checking to make sure all headers are valid
	rows, err := loadCsv.GetCSV(file)
	if err != nil {
		fmt.Println(err)
	}

	// Split the CSV into as many parts as needed, capping at the given file size limit
	isTooLarge, err := fixCsvSize.Check(file)
	if err != nil {
		fmt.Println(err)
	}

	if isTooLarge {
		fixCsvSize.Split("resizedCsv", rows[0], rows)
	}

	fmt.Println("ok")
}
