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
	"fmt"
	"strconv"
)

// IntOpts is
type IntOpts struct {
	Required     bool
	DefaultKey   string
	DefaultValue int
	Help         string // short help message if required and not found
}

// AsInt is retrieves an integer value from the environment.
// Returns an error if the value is missing or invalid.
func AsInt(key string, opts IntOpts) (int, error) {
	_, val, err := GetInt(key, opts)
	return val, err
}

// GetInt searches the environment for a list of keys.
// Returns the matching key's name and value.
// Returns an error if the key is not set or the value is invalid.
func GetInt(key string, opts IntOpts, altKeys ...string) (keyFound string, val int, err error) {
	keys := []string{key}
	if len(altKeys) != 0 {
		keys = append(keys, altKeys...)
	}
	key, sval, err := SearchEnv(keys...)
	if err != nil {
		if err == IsBlank && !opts.Required {
			return opts.DefaultKey, opts.DefaultValue, nil
		}
		return key, opts.DefaultValue, err
	}

	if val, err = strconv.Atoi(sval); err != nil || sval != fmt.Sprintf("%d", val) {
		return key, opts.DefaultValue, InvalidInteger
	}

	return key, val, nil
}
