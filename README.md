# Probexss
This simpe script in Go automates testing for Blind XSS vulnerabilities by injecting payloads into query parameters of provided URLs. It supports GET, POST, and PUT requests while respecting a delay between requests.
### Features
  - Reads URLs from a file.
  - Modifies all query parameters with XSS Payload.
  - Send requests using different HTTP methods.
  - Modifies User-Agent.
  - Mofifies delay per request for avoid detection.
## Installation
`go build -o probexss main.go`
## Use
**The file must be .txt**

`./probexss <file.txt>`

**If you want automatize the script in your VPS and go to sleep you can execute the next command:**

`nohup ./probexss <file.txt> > out.log 2>&1 &`

***<sub>Remember you can paste more xss payloads in the string variable to increase the possibilities :) </sub>***
