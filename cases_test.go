package main

var exampleUrl = "http://example.com"
var exampleBody = []byte("<!doctype html>\n<html>\n<head>\n    <title>Example Domain</title>\n\n    <meta charset=\"utf-8\" />\n    <meta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    <style type=\"text/css\">\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: -apple-system, system-ui, BlinkMacSystemFont, \"Segoe UI\", \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 2em;\n        background-color: #fdfdff;\n        border-radius: 0.5em;\n        box-shadow: 2px 3px 7px 2px rgba(0,0,0,0.02);\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        div {\n            margin: 0 auto;\n            width: auto;\n        }\n    }\n    </style>    \n</head>\n\n<body>\n<div>\n    <h1>Example Domain</h1>\n    <p>This domain is for use in illustrative examples in documents. You may use this\n    domain in literature without prior coordination or asking for permission.</p>\n    <p><a href=\"https://www.iana.org/domains/example\">More information...</a></p>\n</div>\n</body>\n</html>\n")
var exampleUrlInfo = URLInfo{
	Url:         "http://example.com",
	Title:       "Example Domain",
	HtmlVersion: "HTML5",
	Headings: []Heading{
		{Importance: 1, Count: 1},
		{Importance: 2, Count: 0},
		{Importance: 3, Count: 0},
		{Importance: 4, Count: 0},
		{Importance: 5, Count: 0},
		{Importance: 6, Count: 0},
	},
	Links: []Link{
		{
			Link:       "https://www.iana.org/domains/example",
			Internal:   false,
			Accessible: false,
		},
	},
	Login: false,
}

var AssembleResponseCases = []struct {
	u        string
	body     []byte
	expected URLInfo
}{
	{
		u:        exampleUrl,
		body:     exampleBody,
		expected: exampleUrlInfo,
	},
}

var QueryURLCases = []struct {
	u        string
	expected []byte
}{
	{
		u:        exampleUrl,
		expected: exampleBody,
	},
}

var ExtractTagsCases = []struct {
	tag      string
	body     []byte
	expected []string
}{
	{
		tag:      "a",
		body:     exampleBody,
		expected: []string{"<a href=\"https://www.iana.org/domains/example\">More information...</a>"},
	},
}

var CheckHTMLVersionCases = []struct {
	body     []byte
	expected string
}{
	{
		body:     exampleBody,
		expected: "HTML5",
	},
}

var CheckForLoginCases = []struct {
	body     []byte
	expected bool
}{
	{
		body:     exampleBody,
		expected: false,
	},
}

var ParseLinkCases = []struct {
	u        string
	link     string
	expected string
	internal bool
}{
	{
		u:        exampleUrl,
		link:     "<a href=\"https://www.iana.org/domains/example\">More information...</a>",
		expected: "https://www.iana.org/domains/example",
		internal: false,
	},
}

var ParseTagCases = []struct {
	tag      string
	expected string
}{
	{
		tag:      "<title>Test Title</title>",
		expected: "Test Title",
	},
}
