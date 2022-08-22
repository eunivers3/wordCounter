package service

import (
	"bufio"
	"log"
	"net/http"
)

type Service struct {
	parser Parser
}

type Parser interface {
	Preprocess(word string) string
	//SortByWordCount(words map[string]int) map[string]int
	//ProcessHTML()
}

func New(parser Parser) *Service {
	return &Service{
		parser: parser,
	}
}

func (s *Service) CountWords(url string) (map[string]int, error) {
	// load content
	res, err := http.Get(url)
	if err != nil {
		log.Println("failed to get url response")
		return map[string]int{}, err
	}
	defer res.Body.Close()
	//TODO: add html parsing to filter text from body content

	words := make(map[string]int)

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords) // split to chunks
	for scanner.Scan() {
		word := scanner.Text()
		word = s.parser.Preprocess(word) // clean word
		if word != "" {
			words[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("failed to get words: %v", err)
		return nil, err
	}

	// TODO: add sorting ability
	//if sortByCount {
	//	return s.parser.SortByWordCount(words), nil
	//}

	return words, nil
}

// TODO: add pagination
//var data []map[string]int
//
//func GetPageData(page int, itemsPerPage int) []map[string]int {
//	start := (page - 1) * itemsPerPage
//	stop := start + itemsPerPage
//
//	if start > len(data) {
//		return nil
//	}
//
//	if stop > len(data) {
//		stop = len(data)
//	}
//
//	return data[start:stop]
//}
