package problem1

import (
	"encoding/csv"
	"os"
)

type Record struct {
	question 	string
	answer	 	string
}

func readCsv(file string) []*Record {
	fileStream, _ := os.Open(file)
	reader := csv.NewReader(fileStream)
	rows, _ := reader.ReadAll()
	//length := len(rows)
	//csvRecord := [length]Record{}
	var csvRecords []*Record
	for i := 0; i < len(rows) ; i++ {
		row := rows[i]
		record := &Record{row[0], row[1]}
		csvRecords = append(csvRecords, record)
	}
	return csvRecords;
}
