// Package tvdb provides a Go client for the TVDB API v4.
//
// Usage:
//
//	client := tvdb.New("your-api-key")
//	if err := client.Login(context.Background()); err != nil {
//		log.Fatal(err)
//	}
//
//	series, err := client.GetSeriesBase(context.Background(), 12345)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(series.Name)
package tvdb

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultBaseURL is the default TVDB API v4 base URL.
	DefaultBaseURL = "https://api4.thetvdb.com/v4"

	// DefaultTimeout is the default HTTP client timeout.
	DefaultTimeout = 30 * time.Second

	// APIVersion is the TVDB API version this client was built for.
	APIVersion = "4.7.10"
)

// Client is a TVDB API client.
type Client struct {
	httpClient *http.Client
	baseURL    string
	token      string
	apiKey     string
	pin        string
}

// Option is a functional option for configuring the Client.
type Option func(*Client)

// New creates a new TVDB API client with the given API key.
func New(apiKey string, opts ...Option) *Client {
	c := &Client{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		baseURL: DefaultBaseURL,
		apiKey:  apiKey,
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

// WithBaseURL sets a custom base URL.
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithPin sets the subscriber PIN for user-supported API keys.
func WithPin(pin string) Option {
	return func(c *Client) {
		c.pin = pin
	}
}

// WithToken sets a pre-existing bearer token (skips login).
func WithToken(token string) Option {
	return func(c *Client) {
		c.token = token
	}
}

// Login authenticates with the TVDB API and stores the bearer token.
// The token is valid for 1 month.
func (c *Client) Login(ctx context.Context) error {
	body := map[string]string{
		"apikey": c.apiKey,
	}
	if c.pin != "" {
		body["pin"] = c.pin
	}

	var response struct {
		Status string `json:"status"`
		Data   struct {
			Token string `json:"token"`
		} `json:"data"`
	}

	if err := c.post(ctx, "/login", body, &response); err != nil {
		return fmt.Errorf("login failed: %w", err)
	}

	c.token = response.Data.Token
	return nil
}

// Token returns the current bearer token.
func (c *Client) Token() string {
	return c.token
}

// get performs a GET request to the given path.
func (c *Client) get(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.request(ctx, http.MethodGet, path, query, nil, result)
}

// post performs a POST request to the given path with the given body.
func (c *Client) post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.request(ctx, http.MethodPost, path, nil, body, result)
}

// request performs an HTTP request.
func (c *Client) request(ctx context.Context, method, path string, query url.Values, body interface{}, result interface{}) error {
	u := c.baseURL + path
	if len(query) > 0 {
		u += "?" + query.Encode()
	}

	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonBody)
	}

	req, err := http.NewRequestWithContext(ctx, method, u, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(respBody),
		}
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}
