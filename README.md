# H24 Assignment

H24 Assignment is a web app which takes a url and displays information about the document, as per the assignment requirements detailed below.

Build the docker image

    docker build -t h24-assignment .

Run the server on port 8000

    docker run --rm -p 8000:8000 h24-assignment

## Assumptions

- Inaccessible links need to be checked
- An ideal solution only uses the standard library
- The complexity of the project should refelct the complexity of the requirements

## Notes

The primary concern of the app is how to handle link checking. This process cannot be handled in the scope of a normal http request because there is no limit to the time it might take. The user should also be delivered _some_ result once the inital url is checked. In order to handle this on the backend, redis (or similar) could be implemented, however the best solution to this problem is to handle the iterations on the frontend.

## Limitations

- Only works on server-side rendered pages, SPAs would require a web driver or JS interpereter
- No security or rate limiting

## TODO

- Improve Handling of Errors (e.g. rate limiting from external sites breaks assemble)
- Extend Test Coverage
- Clean up links - remove duplicates and non-url
- Deployment
- Implement web toolkit (e.g gorilla)
- Logging

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
2. Provide a README.md in GitHub repository with instructions on how to run the application
3. Send us the GitHub’s link to the repository

Good luck!
