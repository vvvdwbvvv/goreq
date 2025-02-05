package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

// DoRequest performs an HTTP request with the given method, URL, body and headers
func DoRequest(opts RequestOptions) ([]byte, error) {
	// Parse the URL
	parsedURL, err := url.Parse(opts.URL)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest(opts.Method, parsedURL.String(), opts.Body)
	if err != nil {
		return nil, err
	}

	// Set headers
	for key, value := range opts.Headers {
		req.Header.Set(key, value)
	}

	// Set default headers if not provided
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json")
	}
	if req.Header.Get("User-Agent") == "" {
		req.Header.Set("User-Agent", "greq/1.0")
	}
	if opts.Body != nil && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Configure client
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Handle redirects
	if !opts.FollowRedirect {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
	}

	// Configure TLS
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: opts.Insecure,
		},
	}

	// Configure proxy
	if opts.Proxy != "" {
		proxyURL, err := url.Parse(opts.Proxy)
		if err != nil {
			return nil, fmt.Errorf("invalid proxy URL: %v", err)
		}
		transport.Proxy = http.ProxyURL(proxyURL)
	}

	client.Transport = transport

	// Set basic auth
	if opts.BasicAuth != "" {
		parts := strings.SplitN(opts.BasicAuth, ":", 2)
		if len(parts) == 2 {
			req.SetBasicAuth(parts[0], parts[1])
		}
	}

	// Print verbose info
	if opts.Verbose {
		printVerboseRequest(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Print response info unless silent
	if !opts.Silent {
		printResponseInfo(resp, opts.IncludeHeaders)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Write to file if specified
	if opts.OutputFile != "" {
		if err := os.WriteFile(opts.OutputFile, body, 0644); err != nil {
			return nil, fmt.Errorf("failed to write to file: %v", err)
		}
	}

	return body, nil
}

func printVerboseRequest(req *http.Request) {
	color.New(color.FgYellow).Printf("> %s %s %s\n", req.Method, req.URL.Path, req.Proto)
	for k, v := range req.Header {
		color.New(color.FgYellow).Printf("> %s: %s\n", k, v[0])
	}
	color.New(color.FgYellow).Println(">")
}

func printResponseInfo(resp *http.Response, includeHeaders bool) {
	color.New(color.FgBlue, color.Bold).Printf("%s %s\n", resp.Proto, resp.Status)
	if includeHeaders {
		color.New(color.FgBlue).Println("Headers:")
		for k, v := range resp.Header {
			color.New(color.FgBlue).Printf("  %s: %s\n", k, v[0])
		}
		fmt.Println()
	}
}

// BuildURLWithQuery combines base URL with query parameters
func BuildURLWithQuery(base string, params map[string]string) (string, error) {
	u, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
