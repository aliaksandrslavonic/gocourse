package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type TestStruct struct {
	Column1 string
	Column2 string
	Column3 string
}

func main() {

	lines, err := ReadCsv("test.csv")
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := TestStruct{
			Column1: line[0],
			Column2: line[1],
			Column3: line[2],
		}
		fmt.Println(data.Column1 + " " + data.Column2 + " " + data.Column3)
	}
}

func ReadCsv(filename string) ([][]string, error) {

	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
