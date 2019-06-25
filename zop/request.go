package zop

// Request defines the request for ZOP express APIs.
type Request struct {
	URL    string
	Params map[string]string
}
