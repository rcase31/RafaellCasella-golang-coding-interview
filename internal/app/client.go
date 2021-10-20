package app

import (
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

const (
	baseUrl = "https://api.ers.usda.gov/data"

	statesUri = "/arms/state"
)

var (
	defaultQueryParams = map[string]string{
		"page_no": "1",
		"limit":   "20",
		"sort":    "name",
		"order":   "asc",
		"random":  strconv.FormatInt(time.Now().Unix(), 10),
	}
)

func FetchStates() ([]byte, error) {
	resp, err := getClient().R().Get(statesUri)
	if err != nil {
		return []byte{}, err
	}

	return resp.Body(), nil
}

func getClient() *resty.Client {
	client := resty.New()
	client.SetHostURL(baseUrl)
	client.SetHeader("Accept", "application/json")
	client.SetQueryParams(defaultQueryParams)
	client.SetError(DefaultError{"An error occurred"})
	return client
}

type DefaultError struct {
	Message string
}
