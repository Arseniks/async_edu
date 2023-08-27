package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func getRequest(requestedUrl string) (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request with ctx: %w", err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to perform http request: %w", err)
	}

	return res, nil
}

func main() {
	request, err := getRequest("https://www.joom.ru/ru")
	if err != nil {
		return
	}
	fmt.Println(request)
}
