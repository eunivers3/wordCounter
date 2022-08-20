package google

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/eunicebjm/gc/pkg/models"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	baseRequestURL   = "https://maps.googleapis.com/maps/api/geocode"
	JSONOutputFormat = "json"
	AddressParameter = "address"
	APIKeyParameter  = "key"
)

type Client interface {
	GeocodeOne(address string) (models.Result, error)
}

type GeocoderClient struct {
	apiKey string
}

func NewGeocoderClient(apiKey string) (Client, error) {
	if apiKey == "" {
		return nil, errors.New("missing argument: apiKey")
	}

	return &GeocoderClient{
		apiKey: apiKey,
	}, nil
}

func (g GeocoderClient) GeocodeOne(address string) (models.Result, error) {
	parameters := encodeParams(models.Parameters{
		Address: address,
		APIKey:  g.apiKey,
	})
	url := fmt.Sprintf("%s/%s?%s", baseRequestURL, JSONOutputFormat, parameters)

	resp, err := http.Get(url)
	if err != nil {
		return models.Result{}, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Result{}, err
	}

	results, err := toResults(body)
	if err != nil {
		fmt.Println("error_unmarshalling_json")
		return models.Result{}, err
	}

	return results, nil
}

func encodeParams(p models.Parameters) string {
	params := url.Values{}
	if p.Address != "" {
		params.Add(AddressParameter, p.Address)
	}
	if p.APIKey != "" {
		params.Add(APIKeyParameter, p.APIKey)
	}

	return params.Encode()
}

func toResults(content []byte) (models.Result, error) {
	var resp models.Response
	if err := json.Unmarshal(content, &resp); err != nil {
		return models.Result{}, err
	}

	if len(resp.Results) == 0 {
		return models.Result{}, errors.New("result_not_found")
	}

	return resp.Results[0], nil
}
