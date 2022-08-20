package service

import (
	"encoding/json"
	"errors"
	"fmt"
	google "github.com/eunicebjm/gc/internal/google/geocoder"
	"github.com/eunicebjm/gc/pkg/models"
	"log"
	"os"
)

var (
	ErrInvalidParam = errors.New("invalid_parameter")
)

// Geocoder is the service business logic interface.
type Geocoder interface {
	GeocodeOne(address string) (models.Result, error)
	GeocodeBatch(req models.Request) (models.Response, error)
}

type Service struct {
	GoogleClient google.Client
}

func NewService(GoogleClient google.Client) (*Service, error) {
	if GoogleClient == nil {
		return nil, fmt.Errorf("%w: GoogleClient", ErrInvalidParam)
	}

	return &Service{
		GoogleClient: GoogleClient,
	}, nil
}

func (s Service) GeocodeOne(address string) (models.Result, error) {
	res, err := s.GoogleClient.GeocodeOne(address)
	if err != nil {
		return models.Result{}, err
	}

	err = s.saveResult(res)
	if err != nil {
		fmt.Println("error saving result")
	}

	return res, nil
}

// save to db
func (s Service) saveResult(res models.Result) error {
	outPath := "test-output.txt"
	data, _ := json.Marshal(res)
	err := os.WriteFile(outPath, data, 0644)
	if err != nil {
		log.Println(err, "failed to write to file")
	}
	return nil
}

// TODO: implement me
func (s Service) GeocodeBatch(req models.Request) (models.Response, error) {
	return models.Response{}, nil
}
