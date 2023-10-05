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

func NewTransaction(direction, year, date, weekday, country, commodity, transport_mode, measure, value, cumulative string) *transaction {

	t := transaction{
		direction: direction, year: year, date: date, weekday: weekday, country: country,
		commodity: commodity, transport_mode: transport_mode, measure: measure, value: value,
		cumulative: cumulative,
	}

	return &t
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func TransactionAverage(commodity, country, weekday string, transactions map[int]*transaction) float64 {

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

func LogExports(commodity string, transactions map[int]*transaction) string {

	exportsByLocation := map[string]int{}
	var biggestExporter string
	var biggestValue int

	for _, currentTransaction := range transactions {

		if currentTransaction.commodity == commodity && currentTransaction.country != "All" {
			temp, _ := strconv.Atoi(currentTransaction.value)
			exportsByLocation[currentTransaction.country] += temp
		}
	}

	for location, totalValue := range exportsByLocation {

		if totalValue > biggestValue {
			biggestExporter = location
			biggestValue = totalValue
		}
	}

	return biggestExporter
}

func main() {

	file, err := os.Open("transaction_registry.csv")
	Check(err)
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	Check(err)

	csvMap := map[int]*transaction{}

	for id := 0; ; id++ {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		csvMap[id] = NewTransaction(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9])
	}

	fmt.Println(csvMap[0])

}
