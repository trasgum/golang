package main

/*
 Create a program that will read in a quiz provided via a CSV file (more details below) and will then give
 the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
 Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

 The CSV file should default to problems.csv, but the user should be able to customize the filename via a flag.

 You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

 At the end of the quiz the program should output the total number of questions correct and how many questions there
 were in total. Questions given invalid answers are considered incorrect.
*/

import (
	"fmt"
	"github.com/docopt/docopt-go"
	"encoding/csv"
	"os"
	"bufio"
	"strings"
)

func main() {
	usage := `Gophercises quiz-1.

	Usage:
	quiz [options]

	Examples:
	quiz -h | --help | -v | --version
	quiz --csv ./problems.csv
	quiz --csv ./problems.csv --limit 60

	Options:
	-h --help            show this help message and exit
	-v --version         show version and exit
	--csv=FILE           a csv file in the format 'question,answer' [default: problems.csv]
	--limit=TIME         bthe time limit in seconds for the quiz [default: 30]
`
	arguments, _ := docopt.Parse(usage, nil, true, "0.1.0-SNAPSHOT", true)

	f, err := os.Open(arguments["--csv"].(string))
	if err != nil {
		fmt.Printf("> Failed reading file! %v\n", err)
		os.Exit(1)
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
        	fmt.Printf("> Failed parsing csv! %v\n", err)
		os.Exit(2)
    	}

	reader := bufio.NewReader(os.Stdin)
	answer_ok := 0
	answer_ko := 0

	for i, line := range lines {
		quiz := line[0]
		answer := line[1]

		fmt.Printf("[#%d] %s: ", i+1, quiz)
		text, _ := reader.ReadString('\n')
		if answer == strings.TrimSpace(text) {
			//fmt.Println("Correct")
			answer_ok += 1
		} else {
			//fmt.Println("Incorrect")
			answer_ko += 1
		}
	}
	fmt.Printf("You scored %d out of %d correct questions", answer_ok, answer_ok+answer_ko)
}

