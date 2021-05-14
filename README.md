# H24 Assignment

Very work in progress, you can currently spin the server up with:

    go run main.go api.go services.go types.go

## Assumptions

- Inaccessible links need to be tested
- An ideal solution uses the standard library
- The link checking is handled by the client to some degree, rather than putting this on the backend and implementing a queuing service like redis

## Limitations

- Only works on server-side rendered pages, SPAs would require a web driver or JS interpereter
- No security or rate limiting

## TODO

- Implement Docker & CI
- Improve Handling of Errors (e.g. rate limiting from external sites breaks assemble)
- Extend Test Coverage
- Clean up links - remove duplicates and non-url
- Handle encoding issue with title (&amp;)
- Handle missing http:// at start of url and subdomains

---

_Original Task Requirements:_

## Golang Technical Challenge

### Task description

Create a web application which takes a website URL as an input and provides general information
about the contents of the page:

- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

Submit the assignment

1. Create a public GitHub repository and commit code in it
2. Provide a README.md in GitHub repository with instructions on how to run the
   application
3. Send us the GitHub’s link to the repository
   Good luck!
