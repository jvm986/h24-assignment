package main

type URLInfo struct {
	Url         string    `json:"url"`
	Title       string    `json:"title"`
	HtmlVersion string    `json:"htmlVersion"`
	Headings    []Heading `json:"headings"`
	Links       []Link    `json:"links"`
	Login       bool      `json:"login"`
}

type Heading struct {
	Importance int `json:"importance"`
	Count      int `json:"count"`
}

type Link struct {
	Link       string `json:"link"`
	Internal   bool   `json:"internal"`
	Accessible bool   `json:"accessible"`
}

// AddHeading appends a Heading to the URLInfo Slice
func (u *URLInfo) AddHeading(h Heading) []Heading {
	u.Headings = append(u.Headings, h)
	return u.Headings
}

// AddLink appends a Link to the URLInfo Slice
func (u *URLInfo) AddLink(l string, internal bool) []Link {
	u.Links = append(u.Links, Link{l, internal, false})
	return u.Links
}
