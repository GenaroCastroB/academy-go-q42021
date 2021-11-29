package common

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type CsvReader struct{}

func NewCsvReader() CsvReader {
	return CsvReader{}
}

func (csvR CsvReader) ReadCsvFile(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer file.Close()

	records, err := readFile(file)
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
		return nil, err
	}

	return records, nil
}

func readFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, err
}
