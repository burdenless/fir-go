package firGo

import (
	"bytes"
	"encoding/json"
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

	// Client Methods
	Incidents IncidentInterface
	Artifacts ArtifactInterface
	Users     UsersInterface
}

// NewFIRClient returns a new FIR API client.
func NewFIRClient(baseHost string, token string) (c *Client) {
	client := http.DefaultClient

	baseURL, _ := url.Parse(baseHost + "/api")
	c = &Client{client: client, BaseURL: baseURL, UserAgent: userAgent, Token: "Token " + token}
	c.Incidents = &IncidentServiceObj{client: c}
	c.Artifacts = &ArtifactServiceObj{client: c}
	c.Users = &UserServiceObj{client: c}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method string, path string, params interface{}) (*http.Request, error) {
	base := c.BaseURL.String()
	fullURL := fmt.Sprintf("%s%s", base, path)

	buf := new(bytes.Buffer)
	if params != nil {
		err := json.NewEncoder(buf).Encode(params)
		if err != nil {
			return nil, err
		}
	}

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

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
