package xii

import (
	"fmt"
	"os"
	"strings"
)

// StringOpts is
type StringOpts struct {
	Required     bool
	DefaultValue string
	Help         string // short help message if required and not found
}

// AsString retrieves a string from the environment.
// Returns an error if the string is not set.
func AsString(key string, opts StringOpts) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return "", fmt.Errorf("%s: must be exported", key)
			}
			return "", fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return "", fmt.Errorf("%s: must be set", key)
	} else if val != trimmedVal {
		return "", fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	return val, nil
}
