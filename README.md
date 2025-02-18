# Probebxss
This simpe script in Go automates testing for Blind XSS vulnerabilities by injecting payloads into query parameters of provided URLs. It supports GET, POST, and PUT requests while respecting a delay between requests.
### Features
  - Reads URLs from a file.
  - Modifies all query parameters with BXSS Payload.
  - Send requests using different HTTP methods.
  - Modifies User-Agent.
  - Mofifies delay per request for avoid detection.
## Installation
`go build -o probebxss main.go`
## Use
**The file must be .txt**

`./probebxss <file.txt>`

**If you want automatize the script in your VPS and go to sleep you can execute the next command:**

`nohup ./probebxss <file.txt> > out.log 2>&1 &`

***<sub>Remember you can paste more bxss payloads in the string variable to increase the possibilities :) </sub>***
