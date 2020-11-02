package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	payload := [][]string{
		{"test1", "test2", "test3"},
		{"somedata", "someresults", "some"},
		{"bla", "blbbla", "blablabla"},
	}

	csvFile, err := os.Create("test.csv")

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)

	for _, empRow := range payload {
		_ = csvWriter.Write(empRow)
	}

	csvWriter.Flush()

}
