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

import "fmt"

// StringOpts is
type StringOpts struct {
	Required     bool
	DefaultKey   string
	DefaultValue string
	Help         string // short help message if required and not found
}

// AsString retrieves a string from the environment.
// Returns an error if the string is not set.
func AsString(key string, opts StringOpts) (string, error) {
	_, val, err := GetString(opts, key)
	return val, err
}

// GetString searches the environment for a list of keys.
// Returns the matching key's name and value.
// Returns an error if the key is not set or the value is invalid.
func GetString(opts StringOpts, keys ...string) (keyFound string, val string, err error) {
	fmt.Printf("[xii] string keys %q\n", keys)

	key, sval, err := SearchEnv(keys...)
	if err != nil {
		if err == NotExported && !opts.Required {
			return opts.DefaultKey, opts.DefaultValue, nil
		}
		return key, sval, err
	}
	return key, sval, nil
}
