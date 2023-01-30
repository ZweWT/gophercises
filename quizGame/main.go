package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	filePtr := flag.String("csv", "problems.csv", "file path to quiz csv")
	timePtr := flag.Int("limit", 30, "the time limit for quiz")
	flag.Parse()

	file, err := os.Open(*filePtr)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	record, err := r.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	correct := 0
	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)

problemloop:
	for i, line := range record {
		fmt.Printf("Problem #%d: %s = ", i, line[0])
		answerChan := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case answer := <-answerChan:
			if answer == line[1] {
				correct++
			}
		case <-timer.C:
			fmt.Println()
			break problemloop
		}

	}

	fmt.Printf("You scored %d out of %d. \n", correct, len(record))
}
