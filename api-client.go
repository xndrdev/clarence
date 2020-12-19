package api_client

import (
	"net/http"
	"fmt"
	"time"
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

func (c *Client) sendRequest(method string, url string, ) error {
	
	// limit := 100
    // page := 1
    // if options != nil {
    //     limit = options.Limit
    //     page = options.Page
    // }

    // req, err := http.NewRequest("GET", fmt.Sprintf("%s/faces?limit=%d&page=%d", c.BaseURL, limit, page), nil)
    // if err != nil {
    //     return nil, err
    // }

    // req = req.WithContext(ctx)

    // res := FacesList{}
    // if err := c.sendRequest(req, &res); err != nil {
    //     return nil, err
    // }

    // return &res, nil
	
	
	
	
	// req.Header.Set("Content-Type", "application/json; charset=utf-8")
	// req.Header.Set("Accept", "application/json; charset=utf-8")
	
	// res, err := c.HTTPClient.Do(req)
    // if err != nil {
    //     return err
    // }

    // defer res.Body.Close()

    // if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
    //     var errRes errorResponse
    //     if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
    //         return errors.New(errRes.Message)
    //     }

    //     return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
    // }

    // fullResponse := successResponse{
    //     Data: v,
    // }
    // if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
    //     return err
    // }

    // return nil
}

type successResponse struct {
	Code int `json:"code"`
	Data string `json:"data"`
}

type errorResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

type Request struct {
	Method string `json:"method"`
	Url string `json:"url"`
	Data interface{}
}