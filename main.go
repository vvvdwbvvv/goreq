package main

import (
	"bytes"
	"flag"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	// Basic options
	requestURL := flag.String("url", "", "HTTP request URL")
	method := flag.String("X", "GET", "HTTP method")
	header := flag.String("H", "", "Custom HTTP header (e.g. 'Authorization: Bearer token')")
	data := flag.String("d", "", "Request body data (for POST/PUT)")
	query := flag.String("q", "", "Query parameters (e.g. 'key1=value1&key2=value2')")

	// Additional curl-like options
	userAgent := flag.String("A", "", "User-Agent string")
	cookie := flag.String("b", "", "Cookie string (e.g. 'name=value')")
	referer := flag.String("e", "", "Referer URL")
	followRedirect := flag.Bool("L", false, "Follow redirects")
	auth := flag.String("u", "", "Basic auth (e.g. 'user:pass')")
	includeHeaders := flag.Bool("i", false, "Include response headers in output")
	proxy := flag.String("x", "", "Proxy URL (e.g. 'http://proxy:8080')")
	insecure := flag.Bool("k", false, "Allow insecure SSL connections")
	silent := flag.Bool("s", false, "Silent mode")
	output := flag.String("o", "", "Write output to file instead of stdout")
	verbose := flag.Bool("v", false, "Verbose output")

	flag.Parse()

	if *requestURL == "" {
		color.Red("Please provide a URL with -url flag")
		flag.Usage()
		os.Exit(1)
	}

	// Load config first
	config := LoadConfig()

	// Build URL with query parameters
	finalURL := *requestURL
	if *query != "" {
		params := parseQueryParams(*query)
		url, err := BuildURLWithQuery(finalURL, params)
		if err != nil {
			color.Red("Error building URL: %v\n", err)
			os.Exit(1)
		}
		finalURL = url
	}

	// Combine headers from config and command line
	headers := make(map[string]string)
	for k, v := range config.DefaultHeaders {
		headers[k] = v
	}

	// Add custom headers
	if *header != "" {
		key, value := parseHeader(*header)
		if key != "" {
			headers[key] = value
		}
	}
	if *userAgent != "" {
		headers["User-Agent"] = *userAgent
	}
	if *referer != "" {
		headers["Referer"] = *referer
	}
	if *cookie != "" {
		headers["Cookie"] = *cookie
	}

	// Prepare request options
	opts := RequestOptions{
		Method:         *method,
		URL:            finalURL,
		Headers:        headers,
		FollowRedirect: *followRedirect,
		BasicAuth:      *auth,
		Proxy:          *proxy,
		Insecure:       *insecure,
		Silent:         *silent,
		Verbose:        *verbose,
		OutputFile:     *output,
		IncludeHeaders: *includeHeaders,
	}

	// Prepare request body if provided
	if *data != "" {
		opts.Body = bytes.NewBuffer([]byte(*data))
	}

	// Make the HTTP request
	resp, err := DoRequest(opts)
	if err != nil {
		color.Red("Error making request: %v\n", err)
		os.Exit(1)
	}

	// Pretty print the response
	PrettyPrintResponse(resp, opts)
}

// parseQueryParams converts query string to map
func parseQueryParams(query string) map[string]string {
	params := make(map[string]string)
	pairs := strings.Split(query, "&")
	for _, pair := range pairs {
		parts := strings.SplitN(pair, "=", 2)
		if len(parts) == 2 {
			params[parts[0]] = parts[1]
		}
	}
	return params
}

// parseHeader splits a header string into key and value
func parseHeader(header string) (string, string) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
