package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

// reading files requires error handling.
// helper function to handle error output

func main() {
	csvFilename := flag.String("file", "problems.csv", "a CSV file with questions and answers")

	flag.Parse()
	quizHandler(readCsv(*csvFilename))
}

func quizHandler(csvData [][]string) {
	counter := 0
	problems := parseQuiz(csvData)
	for i, p := range problems {
		fmt.Printf("Question #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			counter++
		}
	}
	fmt.Printf("You got %d out of %d correct", counter, len(problems))
}

func readCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	errorHandler(err, "Failed to open CSV")

	r := csv.NewReader(file)
	r.Comma = ','
	r.FieldsPerRecord = 2
	data, err := r.ReadAll()
	errorHandler(err, "Failed to parse data")

	return data
}

func parseQuiz(lines [][]string) []problem {
	returnedSlice := make([]problem, len(lines))
	for num, row := range lines {
		returnedSlice[num] = problem{
			q: row[0],
			a: strings.TrimSpace(row[1]),
		}
	}
	return returnedSlice
}

func errorHandler(e error, eMessage string) {
	if e != nil {
		fmt.Printf("%s\n", eMessage)
		os.Exit(1)
	}
}

type problem struct {
	q string
	a string
}
