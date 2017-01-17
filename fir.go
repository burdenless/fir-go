package firGo

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

const (
	libraryVersion = "0.1.0"
	userAgent      = "firGo/" + libraryVersion
	mediaType      = "application/json"
)

// Client manages communication with FIR API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests.
	BaseURL *url.URL

	// User agent for client
	UserAgent string

	// Token for authentication
	Token string
}

// NewFIRClient returns a new FIR API client.
func NewFIRClient(baseHost string, token string) (c *Client) {
	client := http.DefaultClient

	baseURL, _ := url.Parse(baseHost + "/api")
	c = &Client{client: client, BaseURL: baseURL, UserAgent: userAgent, Token: "Token " + token}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string) (*http.Request, error) {
	base := c.BaseURL.String()
	fullURL := fmt.Sprintf("%s%s", base, path)
	fmt.Printf("[*] URL being called: %s\n", fullURL)

	buf := new(bytes.Buffer)
	req, err := http.NewRequest(method, fullURL, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", mediaType)
	req.Header.Add("User-Agent", c.UserAgent)

	return req, err
}

// Do sends an API request and returns the API response. The API response is JSON decoded and stored in the value
// pointed to by v, or returned as an error if an API error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
