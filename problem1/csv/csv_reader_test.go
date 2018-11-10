package csv

import "testing"

func TestShouldReadMultilineCSV(t *testing.T) {
	records, _ := ReadCsv("simple.csv")
	validateCsvFormat(records, 13, t)
}

func TestShouldReadCSVWithSpecialCharacters(t *testing.T) {
	records, _ := ReadCsv("csv_with_special_characters.csv")
	validateCsvFormat(records, 3, t)
}

func TestShouldNotErrorOutEmptyFile(t *testing.T) {
	records, _ := ReadCsv("empty.csv")
	validateCsvFormat(records, 0, t)
}


func TestShouldThrowErrorForInvalidCsv(t *testing.T) {
	_, err := ReadCsv("invalid.csv")
	if (err == nil) {
		t.Errorf("Invalid CSV should throw error.")
	}
}

func validateCsvFormat(records []*Record, numRecords int, t *testing.T) {
	if len(records) != numRecords {
		t.Errorf("Number of records didn't match. Expected: %d, Actual :%d", numRecords, len(records))
	}

	for i:= 0; i < len(records); i++ {
		if(records[i].Question == "") {
			t.Errorf("Record : %d has question undefined", i)
		}
		if(records[i].Answer == "") {
			t.Errorf("Record : %d has answer undefined", i)
		}
	}
}