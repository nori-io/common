package http

import (
	"context"
	"net/http"
	"net/url"

	"github.com/nori-io/nori-common/endpoint"
)

type HTTPClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client struct {
	client HTTPClient
	method string
	url    *url.URL
	enc    EncodeRequest
	dec    DecodeResponse
	before []BeforeFunc
	after  []ClientAfterFunc
}

type ClientOption func(client *Client)

// By default, http.DefaultClient is used.
func SetClient(client HTTPClient) ClientOption {
	return func(c *Client) { c.client = client }
}

func NewClient(
	method string,
	url *url.URL,
	enc EncodeRequest,
	dec DecodeResponse,
	options ...ClientOption,
) *Client {
	c := &Client{
		client: http.DefaultClient,
		method: method,
		url:    url,
		enc:    enc,
		dec:    dec,
		before: []BeforeFunc{},
		after:  []ClientAfterFunc{},
	}
	for _, option := range options {
		option(c)
	}
	return c
}

// Endpoint returns a usable endpoint that invokes the remote endpoint.
func (c Client) Endpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		ctx, cancel := context.WithCancel(ctx)

		var (
			resp *http.Response
			err  error
		)

		req, err := http.NewRequest(c.method, c.url.String(), nil)
		if err != nil {
			cancel()
			return nil, err
		}

		if err = c.enc(ctx, req, request); err != nil {
			cancel()
			return nil, err
		}

		for _, f := range c.before {
			ctx = f(ctx, req)
		}

		resp, err = c.client.Do(req.WithContext(ctx))

		if err != nil {
			cancel()
			return nil, err
		}

		for _, f := range c.after {
			ctx = f(ctx, resp)
		}

		response, err := c.dec(ctx, resp)
		if err != nil {
			return nil, err
		}

		return response, nil
	}
}
