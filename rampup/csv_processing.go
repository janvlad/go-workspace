package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type transaction struct {
	direction, year, date, weekday, country, commodity, transport_mode, measure, value, cumulative string
}

func newTransaction(direction, year, date, weekday, country, commodity, transport_mode, measure, value, cumulative string) *transaction {

	t := transaction{
		direction: direction, year: year, date: date, weekday: weekday, country: country,
		commodity: commodity, transport_mode: transport_mode, measure: measure, value: value,
		cumulative: cumulative,
	}

	return &t
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func transactionAverage(commodity, country, weekday string, transactions map[int]*transaction) float64 {

	var sumOfValues, counter float64

	for _, currentTransaction := range transactions {

		if currentTransaction.country == country && currentTransaction.commodity == commodity && currentTransaction.weekday == weekday {
			temp, _ := strconv.ParseFloat(currentTransaction.value, 64)
			sumOfValues += temp
			counter++
		}
	}

	return (sumOfValues / counter)
}

func main() {

	file, err := os.Open("transaction_registry.csv")
	check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	check(err)

	csvMap := map[int]*transaction{}

	for id := 0; ; id++ {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		csvMap[id] = newTransaction(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9])
	}

	fmt.Println(csvMap[0])

}
