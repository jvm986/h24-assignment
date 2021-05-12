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
func AssembleResponse(url string, body []byte) (URLInfo, error) {
	info := URLInfo{}
	info.Url = url
	// Needs work
	titleTag, _ := ExtractTags("title", body)
	info.Title, _ = ParseTag("title", titleTag[0])
	info.HtmlVersion, _ = CheckHTMLVersion(body)

	for i := 1; i <= 6; i++ {
		r, _ := ExtractTags(fmt.Sprintf("h%d", i), body)
		h := Heading{i, len(r)}
		info.AddHeading(h)
	}
	r, _ := ExtractTags("a", body)
	for _, u := range r {
		l, i, err := ParseLink(url, u)
		if err == nil {
			li := Link{l, i}
			info.AddLink(li)
		}
	}
	info.Login, _ = CheckForLogin(body)
	return info, nil
}

// QueryURL takes a url and returns the HTML document body
func QueryURL(u string) ([]byte, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return []byte{}, errors.New("invalid url")
	}
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	//TODO: investigate errors and handle
	request, _ := http.NewRequest("GET", u, nil)
	request.Header.Set("user-agent", "h24-assignment https://github.com/jvm986/h24-assignment"+
		" - Checks for page contents of any given URL")
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
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
	re := regexp.MustCompile("(?i)^<!DOCTYPE html>")
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
func ParseLink(url string, l string) (string, bool, error) {
	r := regexp.MustCompile("(?i)href=\"(.*?)\"").FindStringSubmatch(l)
	i := regexp.MustCompile(fmt.Sprintf("(?i)^(/|%v)", url)).MatchString(l)
	if len(r) >= 2 {
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
