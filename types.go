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
	Url      string `json:"Url"`
	Internal bool   `json:"internal"`
}

// AddHeading appends a Heading to the URLInfo Slice
func (u *URLInfo) AddHeading(h Heading) []Heading {
	u.Headings = append(u.Headings, h)
	return u.Headings
}

// AddLink appends a Link to the URLInfo Slice
func (u *URLInfo) AddLink(l Link) []Link {
	u.Links = append(u.Links, l)
	return u.Links
}
