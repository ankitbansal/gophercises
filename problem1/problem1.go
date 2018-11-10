package main

import (
	"bufio"
	"gophercises/problem1/csv"
	"fmt"
	"os"
	"strings"
	"flag"
	"math/rand"
	"time"
)

var (
	filePath 	string
	random		bool
	timeOut		int
)

func init() {
	flag.StringVar(&filePath, "filePath", "csv/simple.csv", "Name of csv File from which to read questions")
	flag.BoolVar(&random, "randomize", true, "boolean flag to specify whether to randomize questions. Defaults to true")
	flag.IntVar(&timeOut, "timeout", 50, "time out for quiz in seconds. Default is 50")
	flag.Parse()
}

func main() {
	records, _ := csv.ReadCsv(filePath)
	reader := bufio.NewReader(os.Stdin)
	correctAnswers := 0

	if random {
		rand.Seed(time.Now().UTC().UnixNano())
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}

	timer := time.NewTimer(time.Second * time.Duration(timeOut))
	defer timer.Stop()

	go func() {
		<-timer.C
		fmt.Printf("You got %d answers right out of %d", correctAnswers, len(records))
		os.Exit(0)
	}()

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