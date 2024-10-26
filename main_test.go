package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintVersion(t *testing.T) {
	testCases := map[string]struct {
		version string
		commit  string
		date    string
		want    string
	}{
		"Test 1": {
			version: "1.0.0",
			commit:  "abc123",
			date:    "2022-01-01",
			want:    "genderize 1.0.0, commit abc123, built at 2022-01-01",
		},
		"Test 2": {
			version: "2.0.0",
			commit:  "def456",
			date:    "2022-02-02",
			want:    "genderize 2.0.0, commit def456, built at 2022-02-02",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			version = tc.version
			commit = tc.commit
			date = tc.date
			got := printVersion()
			assert.Equal(t, tc.want, got)
		})
	}
}
