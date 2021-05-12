package main

import "testing"

func TestParseTag(t *testing.T) {
	for _, test := range ParseTagCases {
		if content, _ := ParseTag("title", test.tag); content != test.expected {
			t.Fatalf("FAIL: ParseTag(%q) = %q, want %q", test.tag, content, test.expected)
		}
	}
}

func TestQueryURL(t *testing.T) {
	for _, test := range QueryURLCases {
		if content, _ := QueryURL(test.url); string(content) != string(test.expected) {
			t.Fatalf("FAIL: QueryURL(%q) = %q, want %q", test.url, content, test.expected)
		}
	}
}
