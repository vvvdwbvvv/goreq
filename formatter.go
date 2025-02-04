package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func PrettyPrintResponse(data []byte, opts RequestOptions) {
	// If output is written to file and silent mode is on, don't print anything
	if opts.OutputFile != "" && opts.Silent {
		return
	}

	var parsed interface{}
	if err := json.Unmarshal(data, &parsed); err != nil {
		// Not JSON, print raw
		fmt.Println(string(data))
		return
	}

	// Pretty print JSON with colors
	pretty, err := json.MarshalIndent(parsed, "", "  ")
	if err != nil {
		fmt.Println(string(data))
		return
	}

	// Color different JSON elements
	lines := strings.Split(string(pretty), "\n")
	for _, line := range lines {
		colorizeJSON(line)
	}
}

func colorizeJSON(line string) {
	keyColor := color.New(color.FgGreen)
	stringColor := color.New(color.FgYellow)
	numberColor := color.New(color.FgCyan)

	// Find the colon that separates key and value
	parts := strings.SplitN(line, ":", 2)

	if len(parts) == 2 {
		// Print key
		keyColor.Print(parts[0] + ":")

		value := strings.TrimSpace(parts[1])
		switch {
		case strings.HasPrefix(value, `"`):
			stringColor.Println(value)
		case strings.ContainsAny(value, "0123456789"):
			numberColor.Println(value)
		default:
			fmt.Println(value)
		}
	} else {
		fmt.Println(line)
	}
}

func parseHeader(header string) (string, string) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
}
