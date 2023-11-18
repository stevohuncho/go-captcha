package capsolver

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	apikey  string
	hClient *resty.Client
	*ClientOpts
}

type ClientOpts struct {
	defaultTimeout int
	defaultPolling int
}

type ClientOptFunc func(*ClientOpts)

func DefaultClientOpts() *ClientOpts {
	return &ClientOpts{
		defaultTimeout: 60000,
		defaultPolling: 5000,
	}
}

func TimeoutClientOpt(timeout int) ClientOptFunc {
	return func(opts *ClientOpts) {
		opts.defaultTimeout = timeout
	}
}

func PollingClientOpt(polling int) ClientOptFunc {
	return func(opts *ClientOpts) {
		opts.defaultPolling = polling
	}
}

func NewClient(apikey string, opts ...ClientOptFunc) *Client {
	c := &Client{
		apikey:     apikey,
		hClient:    resty.New().SetBaseURL("https://api.capsolver.com"),
		ClientOpts: DefaultClientOpts(),
	}
	for _, opt := range opts {
		opt(c.ClientOpts)
	}
	return c
}

func (c *Client) DefaultTimeout() int {
	return c.defaultTimeout
}

func (c *Client) DefaultPolling() int {
	return c.defaultPolling
}
