package main

import (
	"bufio"
	"gophercises/problem1/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	records, _ := csv.ReadCsv("csv/simple.csv")
	reader := bufio.NewReader(os.Stdin)
	correctAnswers := 0

	for i:= 0; i < len(records); i++ {
		record := records[i]
		fmt.Printf("Question : %s=", record.Question)
		text, _ := reader.ReadString('\n')
		if strings.Compare(record.Answer, strings.TrimRight(text, "\n")) == 0 {
			correctAnswers = correctAnswers+1;
		}
	}

	fmt.Printf("You got %d answers right out of %d", correctAnswers, len(records))
}