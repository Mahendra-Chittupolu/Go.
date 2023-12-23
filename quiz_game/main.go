package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	time_limit := flag.Int("limit", 30, "time limit for quiz in seconds")
	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines((lines))
	timer := time.NewTimer(time.Duration(*time_limit) * time.Second)
	fmt.Println((problems))
	correct := 0
	for i, p := range problems {
		select {
		case <-timer.C:
			fmt.Printf("Problem #%d : %s =\n", i+1, p.q)
			return
		default:
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == p.a {
				correct++
			}
		}
	}
	fmt.Printf("you have scored %d out of %d", correct, len(problems))

}
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: line[1],
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
