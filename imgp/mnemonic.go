package imgp

import (
	_ "embed"
	"strings"
)

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
