package utils

type CMCClient interface {
	NewClient(apiKey string) *Client
	GetLatestListings() (string, error)
}
