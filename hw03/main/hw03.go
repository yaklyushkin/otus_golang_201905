package main

import (
	"fmt"
	"sort"
	"strings"
)

type WordsCount struct {
	word  string
	count int
}

func Top10words(str string) []WordsCount {
	var splitted []string = strings.Fields(str)
	var counter map[string]int = make(map[string]int)
	for _, word := range splitted {
		counter[word] += 1
	}

	var words_count []WordsCount
	for key, value := range counter {
		words_count = append(words_count, WordsCount{key, value})
	}

	sort.Slice(words_count, func(i, j int) bool {
		if words_count[i].count == words_count[j].count {
			return words_count[i].word < words_count[j].word
		}
		return words_count[i].count > words_count[j].count
	})

	var result []WordsCount
	var result_len int = 10
	if len(words_count) < result_len {
		result_len = len(words_count)
	}
	for _, data := range words_count[:result_len] {
		result = append(result, data)
	}
	return result
}

func main() {
	var test_data []string = []string{
		"6 11 1 10  2 3 6 5 4 7 9 8 12",
		"asd dfg dfg zxc dfs asd",
		"",
	}
	for _, test := range test_data {
		fmt.Println(test)
		answer := Top10words(test)
		for _, word_count := range answer {
			fmt.Printf("%d\t %s\n", word_count.count, word_count.word)
		}
		fmt.Println("---")
	}
}
