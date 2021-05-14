package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

// AssembleResponse generates the response slice for processing by the server
// TODO: Handle Errors -- embed in struct?
func AssembleResponse(u string, body []byte) (URLInfo, error) {
	info := URLInfo{}
	info.Url = u
	// Needs work
	titleTag, _ := ExtractTags("title", body)
	if len(titleTag) > 0 {
		info.Title, _ = ParseTag("title", titleTag[0])
	} else {
		info.Title = "<Page has no title>"
	}
	info.HtmlVersion, _ = CheckHTMLVersion(body)

	for i := 1; i <= 6; i++ {
		r, _ := ExtractTags(fmt.Sprintf("h%d", i), body)
		h := Heading{i, len(r)}
		info.AddHeading(h)
	}
	r, _ := ExtractTags("a", body)
	for _, u := range r {
		l, i, err := ParseLink(u, u)
		if err == nil {
			if i {
				info.AddLink(l, true)
			} else {
				info.AddLink(l, false)
			}
		}
	}
	info.Login, _ = CheckForLogin(body)
	return info, nil
}

// QueryURL takes a url and returns the HTML document body
func QueryURL(u string) ([]byte, error) {
	re := regexp.MustCompile("^(http://|https://)")
	if !re.Match([]byte(u)) {
		u = "https://" + u
	}
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return []byte{}, errors.New("invalid url")
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	request, _ := http.NewRequest("GET", u, nil)
	request.Header.Set("user-agent", "h24-assignment https://github.com/jvm986/h24-assignment"+
		" - Checks for page contents of any given URL")
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, errors.New("error connecting to host")
	}
	defer response.Body.Close()

	// TODO: check for doctype in header
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading HTTP body. ", err)
		return []byte{}, errors.New("error reading response body")
	}
	return body, nil
}

// ExtractTags takes a tag to search for and returns a list of tags
func ExtractTags(tag string, body []byte) ([]string, error) {
	re := regexp.MustCompile(fmt.Sprintf("(?i)<%s( |>)(.|\n)*?</%s>", tag, tag))
	tags := re.FindAllString(string(body), -1)
	if tags == nil {
		return []string{}, errors.New("no tags found")
	}
	return tags, nil
}

// CheckHTMLVersion takes html body and returns the version
// TODO: extend for responsive & do more research
func CheckHTMLVersion(body []byte) (string, error) {
	re := regexp.MustCompile("(?i)<!DOCTYPE html>")
	if re.Match(body) {
		return "HTML5", nil
	}
	return "HTML 4.01", nil
}

// CheckForLogin takes html body and returns true if there is a login form
// TODO: extend to cover more cases
func CheckForLogin(body []byte) (bool, error) {
	re := regexp.MustCompile("(?i)<input.* type=\"password\"")
	return re.Match(body), nil
}

// ParseLink takes a tag and extracts the url
func ParseLink(u string, l string) (string, bool, error) {
	r := regexp.MustCompile(`(?i)href="(.*?)"`).FindStringSubmatch(l)
	if len(r) >= 2 {
		i := regexp.MustCompile(`(?i)^\/`).MatchString(r[1])
		if i {
			r[1] = u + r[1]
		}
		i = regexp.MustCompile(fmt.Sprintf(`(?i)%s`, u)).MatchString(r[1])
		_, err := url.ParseRequestURI(r[1])
		if err != nil {
			return "", false, errors.New("invalid url")
		}
		return r[1], i, nil
	}
	return "", false, errors.New("invalid link")
}

// ParseTag takes a tag and extract the content
func ParseTag(tag string, l string) (string, error) {
	r := regexp.MustCompile(fmt.Sprintf("(?i)<%s(>| )(.*?)</%s>", tag, tag)).FindStringSubmatch(l)
	if len(r) >= 3 {
		return r[2], nil
	}
	return "", errors.New("invalid tag")
}
