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

// Package xii contains helpers for 12 Factor.
package xii

import (
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
			return opts.DefaultValue, NotExported
		}
		return opts.DefaultValue, nil
	}

	trimmedVal := strings.TrimSpace(val)
	if opts.Required && trimmedVal == "" {
		return opts.DefaultValue, IsBlank
	} else if val != trimmedVal {
		return opts.DefaultValue, ExtraSpaces
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
	return opts.DefaultValue, InvalidBoolean
}
