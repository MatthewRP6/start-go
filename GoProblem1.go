package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := os.Open("marmatt.txt")
	check(err)
	s := bufio.NewScanner(dat)
	wordsList := make([]string, 0)
	answer := map[string]int{}

	for s.Scan() {
		line := s.Text()
		words := strings.Split(line, " ")
		wordsList = append(wordsList, words...)
	}

	for _, v := range wordsList {
		_, ok := answer[v]
		if ok {
			answer[v] += 1
		} else {
			answer[v] = 1
		}

	}
	fmt.Println(answer)

}
