package reader

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

type Reader struct{}

func New() (Reader, error) {
	return Reader{}, nil
}

func (r Reader) LoadText(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println("failed to get url response")
		return []string{}, err
	}
	defer res.Body.Close()

	return linesFromReader(res.Body)
}

func (r Reader) LoadText2(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Println("failed to get url response")
		return "", err
	}
	defer res.Body.Close()

	content, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("failed to load url content")
		return "", err
	}
	return string(content), nil
}

func linesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("failed to get lines", err)
		return nil, err
	}

	return lines, nil
}
