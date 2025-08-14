package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var correctResults = 0
	for _, record := range records {
		//fmt.Println("Record:", record)
		var valuesSum = strings.Split(record[0], "+")
		//fmt.Println("Values", valuesSum)
		sum1, err := strconv.Atoi(valuesSum[0])
		if err != nil {
			panic(err)
		}
		sum2, err := strconv.Atoi(valuesSum[1])
		if err != nil {
			panic(err)
		}
		result, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}

		if sum1+sum2 == result {
			correctResults++
		}
	}

	fmt.Println("Correct results", correctResults)

}
