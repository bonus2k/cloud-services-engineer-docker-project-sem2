package main

import (
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://localhost:8081/health")
	if err != nil || resp.StatusCode >= 400 {
		os.Exit(1)
	}
	os.Exit(0)
}
