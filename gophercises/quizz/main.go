package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func main() {
	//Binary info
	csvFilename := flag.String("csv", "problens.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()
	_ = csvFilename

	//File Opening
	csvArg := os.Args[1]
	csvfile, err := os.Open(csvArg)
	if err != nil {
		fmt.Printf("there was an error opening the file %v", err)
		return
	}
	r := csv.NewReader(csvfile)

	//Funky intro
	introduction()

	//questions and answers
	var rights, wrongs int
	csvParser(r, &rights, &wrongs)
	fmt.Printf("TOTAL IS \n%v / %v GOOD ANSWERS", rights, wrongs)
}

func csvParser(r *csv.Reader, rights *int, wrongs *int) {
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("reading error %v", err)
		}
		fmt.Printf("%v ", record[0])
		a, err := strconv.Atoi(record[1])
		if questionAndAnswer(a) == 1 {
			*rights++
		} else {
			*wrongs++
		}
	}
}

func questionAndAnswer(expected int) int {
	var e int
	fmt.Scanf("%d", &e)
	if e == expected {
		return 1
	} else {
		// fmt.Printf("received %v -- expected %v\n", e, expected)
		return 0
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
