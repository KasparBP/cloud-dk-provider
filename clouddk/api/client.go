package api

import (
	"errors"
	"net/http"
	"net/url"
)

const (
	defaultEndpoint = "https://api.cloud.dk"
	userAgent = "Terraform-cloud-dk-provider"
)

type Client struct {
	baseURL            *url.URL
	httpClient         *http.Client
	token              string
	CloudServerService *CloudServerService
}

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
	client.CloudServerService = &CloudServerService{ client: client }
	return client, nil
}
