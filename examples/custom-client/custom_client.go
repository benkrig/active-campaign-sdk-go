package main

import (
	ac "github.com/benkrig/active-campaign-sdk-go"
	"net/http"
	"os"
)

// If you'd like, you can build your httpClient and avoid passing your token through this package entirely.
func main() {
	client := http.DefaultClient
	rt := WithHeader(client.Transport)
	rt.Set("Api-Token", "your token")
	client.Transport = rt

	baseURL := os.Getenv("YOUR_BASE_URL_KEY")

	a, err := ac.NewClient(
		&ac.ClientOpts{
			HttpClient: client,
			BaseUrl:    baseURL,
		},
	)
	if err != nil {
		panic(err)
	}

	_, _, err = a.Tags.ListAll()
	if err != nil {
		panic(err)
	}
}

// Credit: https://stackoverflow.com/a/51326483/4544386
type MyClient struct {
	http.Header
	rt http.RoundTripper
}

func WithHeader(rt http.RoundTripper) MyClient {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return MyClient{Header: make(http.Header), rt: rt}
}

func (c MyClient) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range c.Header {
		req.Header[k] = v
	}

	return c.rt.RoundTrip(req)
}
