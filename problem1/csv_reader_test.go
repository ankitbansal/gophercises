package problem1

import "testing"

func TestShouldReadMultilineCSV(t *testing.T) {
	records := readCsv("simple.csv")
	if len(records) != 13 {
		t.Errorf("Number of records didn't match. Expected: %d, Actual :%d", 13, len(records))
	}

	for i:= 0; i < len(records); i++ {
		if(records[i].question == "") {
			t.Errorf("Record : %d has question undefined", i)
		}
		if(records[i].answer == "") {
			t.Errorf("Record : %d has answer undefined", i)
		}
	}
}