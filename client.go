package docling

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

type ClientConfig struct {
	APIKey  string
	BaseURL string
}

type ClientOption func(*Client)

func WithHTTPClient(httpCli *http.Client) ClientOption {
	return func(c *Client) {
		c.httpCli = httpCli
	}
}

func WithLogger(logger *slog.Logger) ClientOption {
	return func(c *Client) {
		c.logger = logger
	}
}

func NewClient(cfg ClientConfig, opts ...ClientOption) (*Client, error) {
	u, err := url.Parse(cfg.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}
	c := &Client{
		apiKey:  cfg.APIKey,
		baseURL: u,
		httpCli: &http.Client{},
		logger:  slog.Default(),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

type Client struct {
	apiKey  string
	baseURL *url.URL
	httpCli *http.Client
	logger  *slog.Logger
}

func (c *Client) NewRequest(ctx context.Context, method, path string, in any) (*http.Request, error) {
	var b io.Reader
	if in != nil {
		data, err := json.Marshal(in)
		if err != nil {
			return nil, fmt.Errorf("failed to json marshal in: %w", err)
		}
		b = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.apiURL(path), b)
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}
	if b != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if len(c.apiKey) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, out any) error {
	resp, err := c.httpCli.Do(req)
	if err != nil {
		return fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return HTTPError{
			StatusCode: resp.StatusCode,
			Body:       data,
		}
	}
	if out != nil {
		err = json.Unmarshal(data, out)
		if err != nil {
			return fmt.Errorf("failed to unmarshal response body: %w", err)
		}
	}
	return nil
}

type HTTPError struct {
	StatusCode int
	Body       []byte
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("http error: unexpected status code: %d, body: %s", e.StatusCode, string(e.Body))
}

func (c *Client) apiURL(path string) string {
	return c.baseURL.JoinPath("v1", path).String()
}
