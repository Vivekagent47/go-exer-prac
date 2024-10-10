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
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formate of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFileName))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to pasrse the provided CSV file.")
	}

	problems := parseLines(lines)
	correctAns := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemLoop:
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)

		ansChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			ansChan <- answer
		}()

		select {
		case <-timer.C:
			break problemLoop
		case answer := <-ansChan:
			if answer == prob.a {
				correctAns += 1
			}
		}
	}

	fmt.Printf("\n\nYou Scored %d out of %d\n", correctAns, len(problems))
	defer file.Close()
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
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
