package autify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Config struct {
	Token     string
	ProjectID string
}

type Client struct {
	http     *http.Client
	baseURL  *url.URL
	apiToken string
}

const baseURL = "https://app.autify.com/api/v1/"

// New returns a clinet for connecting to the Autify API
// https://autifyhq.github.io/autify-api/
func New(config Config) *Client {

	baseURL, _ := url.Parse(baseURL)

	c := &Client{
		http:     &http.Client{},
		baseURL:  baseURL,
		apiToken: config.Token,
	}

	return c
}

func (c *Client) get(ctx context.Context, url string, result interface{}) error {
	return c.call(ctx, url, "GET", nil, result)
}

func (c *Client) post(ctx context.Context, url string, postBody interface{}, result interface{}) error {
	jsonParams, err := json.Marshal(postBody)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(jsonParams)

	return c.call(ctx, url, "POST", body, result)
}

func (c *Client) delete(ctx context.Context, url string, result interface{}) error {
	return c.call(ctx, url, "DELETE", nil, result)
}

func (c *Client) call(ctx context.Context, urlStr, method string, body io.Reader, result interface{}) error {
	if !strings.HasSuffix(c.baseURL.Path, "/") {
		return fmt.Errorf("BaseURL must have a trailing slash, but %q does not", c.baseURL)
	}
	url, err := c.baseURL.Parse(urlStr)

	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), nil)

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return c.error(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)

	if err != nil {
		return err
	}

	return nil
}

// Error represents an error created by the autify-go client.
type Error struct {
	// The HTTP status code.
	Status int `json:"status"`
	// A short description of the error.
	Message string `json:"message"`
}

func (e Error) Error() string {
	return e.Message
}

// AutifyError represents an error returned by the autify api.
type AutifyError struct {
	Errors []AutifyErrorDetail `json:"errors,omitempty"`
}

// AutifyErrorDetail represents a message returned by the autify api.
type AutifyErrorDetail struct {
	Message string `json:"message,omitempty"`
}

func (c *Client) error(resp *http.Response) error {
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("autify: HTTP status %d: %s, empty request body", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e AutifyError
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("autify: couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	resError := Error{
		Status: resp.StatusCode,
	}
	if len(e.Errors) > 0 {
		resError.Message = "autify: raw error messages: "
		for _, v := range e.Errors {
			resError.Message += fmt.Sprintf("%s, ", v.Message)
		}
	}

	return resError
}
