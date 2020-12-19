package api_client

import (
	"net/http"
	"fmt"
    "time"
    "strings"
    "io/ioutil"
)

func Hello(name string) string {
	message := fmt.Sprintf("hi, %v. Welcome!", name)
	return message
}

type Client struct {
	BaseURL string
	ClientId string
	Secret string
	HTTPClient *http.Client
}

func NewClient(baseUrl string, clientId string, secret string) *Client {
	return &Client{
		BaseURL: baseUrl,
		ClientId: clientId,
		Secret: secret,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) SendRequest(method string, url string, data interface{}) ([]byte, error) {
    
    path := buildUrl(c.BaseURL, url)
    req, err := http.NewRequest(method, path, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return nil, err
    }

    defer resp.Body.Close()

    return ioutil.ReadAll(resp.Body)
}

func buildUrl(baseUrl string, path string) string {
    var pathBuilder strings.Builder
    
    pathBuilder.WriteString(baseUrl)
    pathBuilder.WriteString("/")
    pathBuilder.WriteString(path)

    return pathBuilder.String()
}