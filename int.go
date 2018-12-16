package xii

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IntOpts is
type IntOpts struct {
	Required     bool
	DefaultValue int
	Help         string // short help message if required and not found
}

// AsInt is retrieves an integer value from the environment.
// Returns an error if the value is missing or invalid.
func AsInt(key string, opts IntOpts) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		if opts.Required {
			if opts.Help == "" {
				return 0, fmt.Errorf("%s: must be exported", key)
			}
			return 0, fmt.Errorf("%s: %s", key, opts.Help)
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return 0, fmt.Errorf("%s must be set", key)
	} else if val != trimmedVal {
		return 0, fmt.Errorf("%s: must not contain leading or trailing spaces", key)
	}

	integer, err := strconv.Atoi(trimmedVal)
	if err != nil || trimmedVal != fmt.Sprintf("%d", integer) {
		return 0, fmt.Errorf("%s: must be a valid integer", key)
	}

	return integer, nil
}
