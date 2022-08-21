package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var (
	ErrInvalidParam = errors.New("invalid_parameter")
)

type Service struct {
	reader Reader
	parser Parser
}

type Reader interface {
	LoadText(url string) ([]string, error)
	LoadText2(url string) (string, error)
}

type Parser interface {
	GetWords(content []string) map[string]int
}

func New(reader Reader, parser Parser) *Service {
	return &Service{
		reader: reader,
		parser: parser,
	}
}

func (s *Service) CountWords(url string) (map[string]int, error) {
	content, err := s.reader.LoadText(url)
	if err != nil {
		return map[string]int{}, errors.New("failed to download text from url")
	}

	// TODO: implement batch
	words := s.parser.GetWords(content)
	if err != nil {
		return map[string]int{}, errors.New("failed to count words")
	}
	//	TODO: save to db
	err = s.saveResult(words)
	if err != nil {
		fmt.Println("error saving result")
	}

	return words, nil
}

// save to db
func (s Service) saveResult(res map[string]int) error {
	outPath := "test-output.txt"
	data, _ := json.Marshal(res)
	err := os.WriteFile(outPath, data, 0644)
	if err != nil {
		log.Println(err, "failed to write to file")
	}
	return nil
}
