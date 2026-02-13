// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package dnspodcn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cases := []struct {
		secretId, secretKey string
	}{
		{"foo", "bar"},
		{"fizz", "buzz"},
	}

	for _, test := range cases {
		c := NewClient(test.secretId, test.secretKey)
		assert.Equal(t, test.secretId, c.SecretId)
		assert.Equal(t, test.secretKey, c.SecretKey)
		assert.Equal(t, apiBaseURL, c.BaseURL)
	}
}
