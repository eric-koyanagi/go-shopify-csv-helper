package main

import (
	"fmt"

	"csvHelper/loadCsv"
)

func main() {
	rows, err := loadCsv.GetCSV("../example.csv")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rows)
}
