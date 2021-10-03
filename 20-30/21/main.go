// https://github.com/gophercises/quiz
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type quiz struct {
	ques string
	ans  string
}

func main() {
	csvFilename := flag.String("filename", "data.csv", "Provide filename for data, default is data.csv'")
	flag.Parse()

	readCsvFile, err := os.ReadFile(*csvFilename)
	if err != nil {
		panic(err)
	}

	csvParse := csv.NewReader(strings.NewReader(string(readCsvFile)))
	csvData, err := csvParse.ReadAll()
	if err != nil {
		panic(err)
	}

	quizzes := make([]quiz, len(csvData))

	for i, line := range csvData {

		quizzes[i] = quiz{ques: line[0], ans: line[1]}
	}

	score := 0

	for i, q := range quizzes {
		fmt.Printf("Quiz #%d: %s = \n", i+1, q.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == q.ans {
			score++
		}
	}

	fmt.Println("Your Score is ", score)

}
