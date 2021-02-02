package zop

import "net/url"

// Request defines the request for ZOP express APIs.
type Request struct {
	URL string

	// if `Body` is provided, use it with `Content-Type: application/json`
	// otherwise, use `Params` with `Content-Type: application/x-www-form-urlencoded`
	Body   string
	Params map[string]string
}

func (r *Request) body() (body string, contentType string) {
	if r.Body != "" {
		body, contentType = r.Body, "application/json"
	} else {
		q := make(url.Values, len(r.Params))
		for k, v := range r.Params {
			q.Set(k, v)
		}
		body, contentType = q.Encode(), "application/x-www-form-urlencoded"
	}
	return
}
