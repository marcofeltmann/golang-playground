package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("got %w, aborting", err)
	}
	fmt.Println("vim-go")
}

type SuperClient struct {
	*http.Client
	token string
}

func (c *SuperClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", "Basic "+c.token)
	return c.Client.Do(req)
}

func run() (err error) {
	req, err := http.NewRequest(http.MethodGet, "https://example.com", nil)
	if err != nil {
		return fmt.Errorf("run(): unable to create request for example.com: %w", err)
	}

	c := &SuperClient{Client: http.DefaultClient, token: "none"}
	resp, err := c.Do(req)
	if err != nil {
		return fmt.Errorf("run(): error getting response for %#v: %w", req, err)
	}
	defer func() { err = resp.Body.Close() }()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("run(): error reading response body from %#v: %w", resp, err)
	}

	fmt.Printf("Got response %q", b)

	return nil
}
