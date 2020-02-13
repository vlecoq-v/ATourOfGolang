package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"math/rand"
	"time"
)

func main() {
	//Binary argument and info
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	shuffleFlag := flag.Bool("shuffle", false, "put this flag if you want to shuffle the questions")
	flag.Parse()

	//File Opening
	csvfile, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("there was an error opening the file %v", err))
	}
	r := csv.NewReader(csvfile)

	//Funky intro
	introduction()

	// read the file and return a slice of problem struct
	problems := csvParser(r)
	if *shuffleFlag {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	//questions and answers from the problems slice
	var rights int
	questionAndAnswers(problems, &rights)
	fmt.Printf("TOTAL IS \n%v / %v GOOD ANSWERS", rights, len(problems))
}

func csvParser(r *csv.Reader) []problem {
		lines, err := r.ReadAll()
		if err != nil {
			exit(fmt.Sprintf("reading error %v", err))
		}
		ret := make([]problem, len(lines))
		for i, line := range lines {
			ret[i] = problem {
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

func questionAndAnswers(problems []problem, rights *int) {
	for i, problem := range problems {
		fmt.Printf("%d %v = ", i + 1, problem.q)
		var userAnswer string
		fmt.Scanf("%s", &userAnswer)
		if userAnswer == problem.a {
			*rights++
		} 
		}
}


func introduction() {
	fmt.Println("This is a math test quizz")
	// time.Sleep(1 * time.Second)
	fmt.Println("REAADY??")
	// time.Sleep(1 * time.Second)
	for i := 0; i < 0; i++ {
		fmt.Println(".")
		time.Sleep(1 * time.Second / 2)
	}
	fmt.Println("GO")
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
