// Package xii contains helpers for 12 Factor.
package xii

import (
	"fmt"
	"os"
	"strings"
)

// BoolOpts specifies if a value is required or has a default value.
type BoolOpts struct {
	Required     bool
	DefaultValue bool
	Help         string // short help message if required and not found
}

// AsBool retrieves a boolean value from the environment.
// Returns an error if the value is missing or invalid.
func AsBool(key string, opts BoolOpts) (bool, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return opts.DefaultValue, fmt.Errorf("%s: must be exported", key)
			}
			return opts.DefaultValue, fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return opts.DefaultValue, fmt.Errorf("%s must be set", key)
	} else if val != trimmedVal {
		return opts.DefaultValue, fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	switch trimmedVal {
	case "false":
		return false, nil
	case "no":
		return false, nil
	case "true":
		return true, nil
	case "yes":
		return true, nil
	}
	return opts.DefaultValue, fmt.Errorf("%s: must be a valid boolean", key)
}
