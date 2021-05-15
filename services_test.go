package main

import (
	"reflect"
	"testing"
)

func TestAssembleResponse(t *testing.T) {
	for _, test := range AssembleResponseCases {
		if content, _ := AssembleResponse(test.u, test.body); !reflect.DeepEqual(content, test.expected) {
			t.Fatalf("FAIL: ParseTag(%q)", test.u)
		}
	}
}
func TestQueryURL(t *testing.T) {
	for _, test := range QueryURLCases {
		if content, _ := QueryURL(test.u); string(content) != string(test.expected) {
			t.Fatalf("FAIL: QueryURL(%q) = %q, want %q", test.u, content, test.expected)
		}
	}
}

func TestExtractTags(t *testing.T) {
	for _, test := range ExtractTagsCases {
		if content, _ := ExtractTags(test.tag, test.body); !reflect.DeepEqual(content, test.expected) {
			t.Fatalf("FAIL: ExtractTags(%q) = \n%q\nwant\n%q", test.tag, content, test.expected)
		}
	}
}

func TestCheckHTMLVersion(t *testing.T) {
	for _, test := range CheckHTMLVersionCases {
		if content, _ := CheckHTMLVersion(test.body); content != test.expected {
			t.Fatalf("FAIL: CheckHTMLVersion() = %q, want %q", content, test.expected)
		}
	}
}

func TestCheckForLogin(t *testing.T) {
	for _, test := range CheckForLoginCases {
		if content, _ := CheckForLogin(test.body); content != test.expected {
			t.Fatalf("FAIL: CheckForLogin() = %t, want %t", content, test.expected)
		}
	}
}

func TestParseLink(t *testing.T) {
	for _, test := range ParseLinkCases {
		if content, internal, _ := ParseLink(test.u, test.link); content != test.expected || internal != test.internal {
			t.Fatalf("FAIL: CheckForLogin(%q, %q) = %q, %t, want %q, %t", test.u, test.link, content, internal, test.expected, test.internal)
		}
	}
}

func TestParseTag(t *testing.T) {
	for _, test := range ParseTagCases {
		if content, _ := ParseTag("title", test.tag); content != test.expected {
			t.Fatalf("FAIL: ParseTag(%q) = %q, want %q", test.tag, content, test.expected)
		}
	}
}
