package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <url>\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("URL is unavailable: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		fmt.Printf("URL is available. Status: %s\n", resp.Status)
		os.Exit(0)
	}

	fmt.Printf("URL returned non-success status: %s\n", resp.Status)
	os.Exit(1)
}
