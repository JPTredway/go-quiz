package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename, timeLimit := parseFlags()
	file, err := getFile(csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	lines, err := getLines(file)
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file: %s\n", *csvFilename))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	runQuiz(timer, problems)
}

func getLines(file *os.File) ([][]string, error) {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	return lines, err
}

func getFile(filename *string) (*os.File, error) {
	file, err := os.Open(*filename)
	return file, err
}

func parseFlags() (*string, *int) {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return csvFilename, timeLimit
}

func runQuiz(timer *time.Timer, problems []problem) {
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d/%d\n", correct, len(problems))
			return
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d/%d\n", correct, len(problems))
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

type problem struct {
	q, a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
