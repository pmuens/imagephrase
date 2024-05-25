package imgn

import (
	_ "embed"
	"fmt"
	"strings"
)

const WordsInMnemonic = 12

//go:embed words.txt
var wordList string
var wordToInt map[string]int
var intToWord map[int]string

func init() {
	// See: https://stackoverflow.com/a/46798310
	splitFn := func(c rune) bool {
		return c == '\n'
	}

	words := strings.FieldsFunc(wordList, splitFn)

	wordToInt = make(map[string]int, len(words))
	intToWord = make(map[int]string, len(words))

	for i, word := range words {
		wordToInt[word] = i
		intToWord[i] = word
	}
}

func WordsToInts(words string) ([]int, error) {
	splitted := strings.Fields(words)
	_ = splitted

	result := make([]int, len(splitted))
	for i, word := range splitted {
		num, ok := wordToInt[word]
		if !ok {
			return nil, fmt.Errorf(`mapping for word "%s" doesn't exist`, word)
		}
		result[i] = num
	}

	return result, nil
}

func IntsToWords(numbers []int) ([]string, error) {
	result := make([]string, len(numbers))
	for i, num := range numbers {
		word, ok := intToWord[num]
		if !ok {
			return nil, fmt.Errorf(`mapping for number "%d" doesn't exist`, num)
		}
		result[i] = word
	}

	return result, nil
}
