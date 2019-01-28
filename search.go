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
	"os"
	"strings"
)

// SearchEnv searches the environment for a list of keys.
// Returns the matching key's name and value.
// Returns an error if the key is not set or the value is invalid.
func SearchEnv(keys ...string) (key, val string, err error) {
	fmt.Printf("[xii] keys %q\n", keys)
	for _, key := range keys {
		fmt.Printf("[xii] key %q\n", key)
		val, ok := os.LookupEnv(key)
		if !ok {
			// not in the environment, so try the next key
			continue
		}
		fmt.Printf("[val] key %q\n", val)

		// if the key is in the environment, it must not be
		// blank or have leading/trailing spaces.
		trimmedVal := strings.TrimSpace(val)
		if len(trimmedVal) == 0 {
			return key, val, IsBlank
		} else if len(trimmedVal) != len(val) {
			return key, val, ExtraSpaces
		}
		return key, val, nil
	}

	return "", "", NotExported
}
