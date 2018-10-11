package accept_test

import (
	"testing"

	accept "github.com/taichi-hagiwara/accept-header"
)

func TestParseContentType(t *testing.T) {
	tests := []string{
		"text/html",
		"text/*",
		"text/html;q=0.1",
		"text/*;q=0.2",
		"*/*",
		"*/*;q=0.2",
	}

	for _, e := range tests {
		c, err := accept.ParseContentType(e)
		if err != nil {
			t.Fatalf("failed to parse: %s, %#v", e, err)
		}
		if c.String() != e {
			t.Fatalf("mismatch result: '%s' != '%s' %#v", e, c.String(), c)
		}
	}
}
