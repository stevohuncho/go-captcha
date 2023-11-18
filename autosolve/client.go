package autosolve

import (
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	apikey    string
	authToken string
	authExp   time.Time
	authWg    *sync.WaitGroup
	hClient   *resty.Client
	tasks     *taskMap
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

func NewClient(apikey string, opts ...ClientOptFunc) *Client {
	c := &Client{
		apikey:     apikey,
		hClient:    resty.New().SetBaseURL("https://api.capsolver.com").SetDebug(true),
		ClientOpts: DefaultClientOpts(),
		tasks:      newTaskMap(),
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
