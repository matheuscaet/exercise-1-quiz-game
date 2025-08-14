package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("problems.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		fmt.Println("Record:", record[0])
	}

}
