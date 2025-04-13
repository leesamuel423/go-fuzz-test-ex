package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		input, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.input)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
