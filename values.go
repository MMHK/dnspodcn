// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"encoding/json"
	"net/url"
	"sort"
	"strings"
)

// Values is a shortcut of map[string]string.
type Values map[string]string

// toValues converts an value to JSON, and then converts JSON to values.
func toValues(v interface{}) (vs Values, err error) {
	data, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &vs)
	return
}

// Get returns the value associated with the given key.
func (vs Values) Get(key string) string {
	v, _ := vs[key]
	return v
}

// Set sets the key to value. It replaces any existing values.
func (vs Values) Set(key, value string) {
	vs[key] = value
}

// Del deletes the value associated with the given key.
func (vs Values) Del(key string) {
	delete(vs, key)
}

// Encode returns
func (vs Values) Encode() string {
	if len(vs) == 0 {
		return ""
	}
	keys := make([]string, 0, len(vs))
	for key := range vs {
		keys = append(keys, key)
	}
	url.Values{}.Encode()
	sort.Strings(keys)
	var buf strings.Builder
	for _, key := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(vs[key]))
	}
	return buf.String()
}
