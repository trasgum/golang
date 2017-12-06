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
	"strings"
	"time"
	"strconv"
)

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func main() {
	usage := `Gophercises quiz-2.

	Usage:
	quiz [options]

	Examples:
	quiz -h | --help | -v | --version
	quiz --csv ./problems.csv
	quiz --csv ./problems.csv --limit 60

	Options:
	-h --help            show this help message and exit
	-v --version         show version and exit
	--csv=<file>         a csv file in the format 'question,answer' [default: problems.csv]
	--limit=<time>       the time limit in seconds for the quiz [default: 30]
`
	arguments, _ := docopt.Parse(usage, nil, true, "0.2.0-SNAPSHOT", true)
	time_limit, err := strconv.Atoi(arguments["--limit"].(string))
	if err != nil {
		fmt.Printf("> Failed reading time limit! %v\n", err)
		os.Exit(1)
	}

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

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(time_limit) * time.Second)
	correct := 0

problemLoop:
	for i, p := range problems {
		fmt.Printf("[#%d] %s: ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <- timer.C:
			fmt.Println()
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d correct questions", correct, len(problems))
}

