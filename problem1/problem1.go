package main

import (
	"bufio"
	"gophercises/problem1/csv"
	"fmt"
	"os"
	"strings"
	"flag"
)

var (
	filePath 	string
	random		bool
)

func init() {
	flag.StringVar(&filePath, "filePath", "csv/simple.csv", "Name of csv File from which to read questions")
	flag.BoolVar(&random, "randomize", true, "boolean flag to specify whether to randomize questions. Defaults to true")
	flag.Parse()
}

func main() {
	records, _ := csv.ReadCsv(filePath)
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