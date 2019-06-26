package zop

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Options sets the options for the `Client`.
type Options func(*options)

type options struct {
	httpClient *http.Client
}

var (
	defaultOptions = options{
		httpClient: http.DefaultClient,
	}
)

// WithHttpClient sets the http client.
func WithHttpClient(client *http.Client) Options {
	return func(o *options) {
		o.httpClient = client
	}
}

// Client defines the client for ZOP express APIs.
type Client struct {
	props *properties
	opts  options
}

// NewClient instantiates a new client.
func NewClient(companyID, key string, opt ...Options) (*Client, error) {
	if companyID == "" {
		return nil, errors.New("missing company id")
	}
	if key == "" {
		return nil, errors.New("missing key")
	}
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	return &Client{
		props: &properties{
			companyID: companyID,
			key:       key,
		},
		opts: opts,
	}, nil
}

// Execute executes the given request.
func (c *Client) Execute(ctx context.Context, req *Request) (string, error) {
	if req.URL == "" {
		return "", errors.New("missing request url")
	}
	q := make(url.Values, len(req.Params))
	for k, v := range req.Params {
		q.Set(k, v)
	}
	body := q.Encode()
	digestBody, err := url.QueryUnescape(body)
	if err != nil {
		return "", err
	}

	httpReq, err := http.NewRequest(http.MethodPost, req.URL, strings.NewReader(q.Encode()))
	if err != nil {
		return "", err
	}
	httpReq.WithContext(ctx)
	httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	httpReq.Header.Add("Content-Type", "charset=utf-8")
	httpReq.Header.Add("x-companyid", c.props.companyID)
	httpReq.Header.Add("x-datadigest", digest(digestBody+c.props.key))

	resp, err := c.opts.httpClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if st := resp.StatusCode; st >= http.StatusBadRequest {
		return "", fmt.Errorf("%d %s", st, http.StatusText(st))
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}
