package newsapi

import (
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const (
	baseUrl    = "https://newsapi.org"
	everything = "/v2/everything"
	top        = "/v2/top-headlines"
	sources    = "/v2/top-headlines/sources"
	apiHeader  = "X-Api-Key"
)

type Client struct {
	client    *http.Client
	apiKey    string
	userAgent string
}

// Create client
func NewClient(api, agent string) *Client {
	c := &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		apiKey:    api,
		userAgent: agent,
	}
	return c
}

// Print Top Headlines
// Print Sources

// -------------------------------------------------------

// Print Everything
func (c *Client) Everything(params *EverythingRequest) {
	base := createUrl(everything)
	setParams(base, params)
	// createRequest
	// makeRequest
}

// Create Request
func (c *Client) createRequest(u *url.URL) *http.Request {
	req, _ := http.NewRequest("GET", u.String(), nil)
	req.Header.Set(apiHeader, c.apiKey)
	return req
}

// Create Base Url
func createUrl(endpoint string) *url.URL {
	base := fmt.Sprintf("%s%s", baseUrl, endpoint)
	u, _ := url.Parse(base)
	return u
}

// Set Params
func setParams(base *url.URL, params interface{}) {
	p := reflect.TypeOf(params)
	fmt.Println(base, p)
}

// Make Request
func makeRequest() {}
