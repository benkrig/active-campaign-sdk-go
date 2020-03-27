package active_campaign

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// httpClient defines an interface for an http.Client implementation so that alternative
// http Clients can be passed in for making requests
type httpClient interface {
	Do(request *http.Request) (response *http.Response, err error)
}

const (
	headerApiToken    = "Api-Token"
	headerContentType = "Content-Type"
)

// A Client manages communication with the Active Campaign API.
type Client struct {
	// HTTP client used to communicate with the API.
	client httpClient

	// Base URL for API requests.
	baseURL *url.URL

	// Token for API requests.
	token string

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// Services used for talking to different parts of the Active Campaign API.
	Contacts *ContactsService
	Tags     *TagsService
}

type service struct {
	client *Client
}

type ClientOpts struct {
	HttpClient httpClient
	BaseUrl    string
	Token      string
}

// NewClient returns a new Active Campaign API client. httpClient is provided to allow a
// custom client in specialized cases.
// If a nil httpClient is provided, a new http.DefaultClient will be used.
func NewClient(opts *ClientOpts) (*Client, error) {
	var httpClient httpClient
	if opts.HttpClient == nil {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := url.Parse(opts.BaseUrl)
	if err != nil {
		return nil, err
	}
	if !strings.HasSuffix(parsedBaseURL.Path, "/") {
		parsedBaseURL.Path += "/"
	}
	if !strings.HasSuffix(parsedBaseURL.Path, "/api/3/") {
		parsedBaseURL.Path += "api/3/"
	}

	c := &Client{
		client:  httpClient,
		baseURL: parsedBaseURL,
		token:   opts.Token,
	}
	c.common.client = c
	c.Contacts = (*ContactsService)(&c.common)
	c.Tags = (*TagsService)(&c.common)
	return c, nil
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	u, err := c.baseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	// URLs should be specified without a preceding slash since baseURL will have the trailing slash
	u.Path = strings.TrimLeft(u.Path, "/")

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set(headerApiToken, c.token)
	req.Header.Set(headerContentType, "application/json")
	return req, nil
}

// Response is a Active Campaign API response. This wraps the standard http.Response
// returned from Active Campaign.
type Response struct {
	*http.Response

	// Future additional fields will go here.
}

func newResponse(r *http.Response) *Response {
	resp := &Response{Response: r}
	return resp
}

// Do sends an API request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned as an error if an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*Response, error) {
	httpResp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	err = CheckResponse(httpResp)
	if err != nil {
		// Even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return newResponse(httpResp), err
	}

	if v != nil {
		// Open a NewDecoder and defer closing the reader only if there is a provided interface to decode to
		defer func() { _ = httpResp.Body.Close() }()
		err = json.NewDecoder(httpResp.Body).Decode(v)
	}

	resp := newResponse(httpResp)

	if v != nil {
		if w, ok := v.(io.Writer); ok {
			_, _ = io.Copy(w, resp.Body)
		} else {
			decErr := json.NewDecoder(resp.Body).Decode(v)
			if decErr == io.EOF {
				decErr = nil // ignore EOF errors caused by empty response body
			}
			if decErr != nil {
				err = decErr
			}
		}
	}

	return resp, err
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if it has a status code outside the 200 range.
// The caller is responsible to analyze the response body.
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	err := fmt.Errorf("Request failed. Please analyze the request body for more details. Status code: %d", r.StatusCode)
	return err
}
