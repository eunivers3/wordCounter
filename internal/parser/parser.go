package parser

import (
	"strings"
	"unicode"
)

type Parser struct{}

func New() (Parser, error) {
	return Parser{}, nil
}

func (p Parser) GetWords(content []string) map[string]int {
	words := make(map[string]int)
	for _, line := range content {
		for k, v := range wordCount(line) {
			words[k] += v
		}
	}
	return words
}

// separate by alphabet start letter?
// for whole words that are separated by dash or apostrophes
func wordCount(line string) map[string]int {
	fields := strings.FieldsFunc(line, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	words := make(map[string]int)
	for _, field := range fields {
		words[strings.ToLower(field)]++
	}
	return words
}

//func preprocess(line string) string {
//	// remove all punctuation
//	// handle doubled dashes
//	// keep apostrophes except
//	// to lower
//}

//func main() {
//	text := "This is a big apple tree. I love big big apple! 42 life--changing iii. you're"
//	fields := strings.FieldsFunc(text, func(r rune) bool {
// returns whole words ignoring punctuation
//return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')
// returns whole words with punctuation
// need to rm everything except  ' and -
//		return !('a' <= r && r <= 'z' || 'A' <= r && r <= 'Z')
//	})
//	words := make(map[string]int)
//	for _, field := range fields {
//		words[strings.ToLower(field)]++
//	}
//	fmt.Println(words)
//}

//"[^\w\d'-\s]+"

//func wordCount(str string) map[string]int {
//	wordList := strings.Fields(str)
//	counts := make(map[string]int)
//	for _, word := range wordList {
//		_, ok := counts[word]
//		if ok {
//			counts[word] += 1
//		} else {
//			counts[word] = 1
//		}
//	}
//	return counts
//}
