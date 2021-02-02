package zop

import "net/url"

// Request defines the request for ZOP express APIs.
type Request struct {
	URL    string
	Params map[string]string
}

func (r *Request) body() (body string, contentType string) {
	q := make(url.Values, len(r.Params))
	for k, v := range r.Params {
		q.Set(k, v)
	}
	body, contentType = q.Encode(), "application/x-www-form-urlencoded"
	return
}
