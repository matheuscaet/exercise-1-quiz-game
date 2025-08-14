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

		var valuesSum = strings.Split(record[0], "+")

		sum1 := getIntValue(valuesSum[0])
		sum2 := getIntValue(valuesSum[1])
		result := getIntValue(record[1])

		if sum1+sum2 == result {
			correctResults++
		}
	}

	fmt.Println("Correct results", correctResults)

}

func getIntValue(s string) int {
	value, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return value
}
