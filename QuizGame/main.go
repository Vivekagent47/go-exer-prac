package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file in the formate of 'question,answer'")
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
	for i, prob := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, prob.q)

		var answer string
		fmt.Scanf("%s\n", &answer)

		if answer == prob.a {
			correctAns += 1
		}
	}

	fmt.Printf("You Scored %d out of %d", correctAns, len(problems))
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
