package client

import (
	"fmt"
	"io"
	"net/http"
)

func New(address string, httpClient *http.Client) *Client {
	return &Client{
		address:    address,
		httpClient: httpClient,
	}
}

type Client struct {
	address    string
	httpClient *http.Client
}

func (c *Client) SetColor(colorHex string) (respBody string, err error) {
	return c.do(fmt.Sprintf("%s/color?color=%s", c.address, colorHex))
}

func (c *Client) SetBrightness(brightness uint16) (respBody string, err error) {
	return c.do(fmt.Sprintf("%s/brightness?brightness=%d", c.address, brightness))
}

func (c *Client) SetMode(mode uint16) (respBody string, err error) {
	return c.do(fmt.Sprintf("%s/mode?mode=%d", c.address, mode))
}

func (c *Client) do(url string) (respBody string, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create new request: %w", err)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to do http: %w", err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}
	return string(bytes), nil
}
