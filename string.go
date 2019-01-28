// xii: a twelve factor helper
//
// Copyright (c) 2019 Michael D Henderson
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package xii

import (
	"os"
	"strings"
)

// StringOpts is
type StringOpts struct {
	Required     bool
	DefaultValue string
	Help         string   // short help message if required and not found
	Alt          []string // alternate keys to search for
}

// AsString retrieves a string from the environment.
// Returns an error if the string is not set.
func AsString(key string, opts StringOpts) (string, error) {
	keys := []string{key}
	if len(opts.Alt) != 0 {
		keys = append(keys, opts.Alt...)
	}

	for _, key := range keys {
		val, ok := os.LookupEnv(key)
		if !ok {
			// not in the environment, so try the next key
			continue
		}

		// if the key is in the environment, it must not be
		// blank or have leading/trailing spaces.
		trimmedVal := strings.TrimSpace(val)
		if len(trimmedVal) == 0 && opts.Required {
			return opts.DefaultValue, IsBlank
		} else if val != trimmedVal {
			return opts.DefaultValue, ExtraSpaces
		}

		return val, nil
	}

	if opts.Required {
		return opts.DefaultValue, NotExported
	}

	return opts.DefaultValue, nil
}
