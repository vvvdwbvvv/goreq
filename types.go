package main

import "io"

// RequestOptions holds all the options for making an HTTP request
type RequestOptions struct {
	Method         string
	URL            string
	Headers        map[string]string
	Body           io.Reader
	FollowRedirect bool
	BasicAuth      string
	Proxy          string
	Insecure       bool
	Silent         bool
	Verbose        bool
	OutputFile     string
	IncludeHeaders bool
}
