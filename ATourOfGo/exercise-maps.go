package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	res := map[string]int{}
	words := strings.Fields(s)
	for i := range words {
		res[words[i]] += 1
	}
	return res
}

func main() {
	wc.Test(WordCount)
}

func printMap(mapu map[string]int) {
	for i := range mapu {
		fmt.Printf("values=%v, occurences=%v\n", i, mapu[i])
	}
}
