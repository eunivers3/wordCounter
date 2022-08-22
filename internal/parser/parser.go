package parser

import (
	"regexp"
	"strings"
)

type Parser struct{}

func New() (Parser, error) {
	return Parser{}, nil
}

// Preprocess removes all non-alphabetical characters
func (p Parser) Preprocess(word string) string {
	//	TODO:
	//	 handle dashed words e.g. life-changing
	//	 handle words with apostrophes e.g. I've, you'll etc.
	r, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		return ""
	}
	return strings.ToLower(r.ReplaceAllString(word, ""))
}

//type pair struct {
//	Key   string
//	Value int
//}

//func (p Parser) SortByWordCount(words map[string]int) map[string]int {
//	var pairs []pair
//	for k, v := range words {
//		pairs = append(pairs, pair{k, v})
//	}
//
//	sort.Slice(words, func(i, j int) bool {
//		return pairs[i].Value > pairs[j].Value
//	})
//
//	return mapToWordCount(pairs)
//}
//
//func mapToWordCount(pairs []pair) map[string]int {
//	words := make(map[string]int, len(pairs))
//	for _, p := range pairs {
//		words[p.Key] = p.Value
//	}
//	return words
//}
