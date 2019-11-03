package api

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	apiTokenHeaderName = "x-api-key"
	defaultEndpoint    = "https://api.cloud.dk"
)

type Client struct {
	baseURL        *url.URL
	httpClient     *http.Client
	token          string
	ClouddkService *ClouddkService
}

// Create a new cloud.dk API client
func NewClient(token string, endpoint string) (*Client, error) {
	if token == "" {
		return nil, errors.New("API Access token missing")
	}
	if endpoint == "" {
		endpoint = defaultEndpoint
	}
	baseUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	client := &Client{
		baseURL:    baseUrl,
		httpClient: http.DefaultClient,
		token:      token,
	}
	client.ClouddkService = &ClouddkService{ c: client }
	return client, nil
}

func (c *Client) newRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL.String() + path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add(apiTokenHeaderName, c.token)
	return req, nil
}

func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	// TODO pagination
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
