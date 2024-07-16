package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	APIkey  string
	BaseURL string
}

func NewClient(apiKey string) *Client {
	return &Client{
		APIkey:  apiKey,
		BaseURL: "https://pro-api.coinmarketcap.com/v1/",
	}

}

func (c *Client) GetLatestListings() (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.BaseURL+"cryptocurrency/listings/latest", nil)
	if err != nil {
		return "", err
	}

	q := url.Values{}

	q.Add("start", "1")
	q.Add("limit", "5000")
	q.Add("convert", "USD")

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", c.APIkey)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}
