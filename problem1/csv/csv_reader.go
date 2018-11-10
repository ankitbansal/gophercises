package csv

import (
	"encoding/csv"
	"os"
	"fmt"
	"errors"
)

type Record struct {
	Question 	string
	Answer	 	string
}

func ReadCsv(file string) ([]*Record, error) {
	var err error
	var csvRecords []*Record

	fileStream, err := os.Open(file)
	if (err != nil) {
		return csvRecords, err
	}

	reader := csv.NewReader(fileStream)
	rows, err := reader.ReadAll()
	if(err != nil) {
		return csvRecords, err
	}

	for i := 0; i < len(rows) ; i++ {
		row := rows[i]
		if (len(row) < 2) {
			err = errors.New(fmt.Sprintf("Row %d is invalid", i))
			return csvRecords, err
		}
		record := &Record{row[0], row[1]}
		csvRecords = append(csvRecords, record)
	}
	return csvRecords, err;
}
