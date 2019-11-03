package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	status, err := healthCheck("https://example.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(status)

}

func healthCheck(URL string) (status int, err error) {
	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), &bytes.Buffer{})
	if err != nil {
		return 0, err
	}

	c := http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
