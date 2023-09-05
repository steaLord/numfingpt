package functions

import (
	"fmt"
	"io"
	"net/http"
)

func FetchData(url string) ([]byte, error) {
	if url == "" {
		return nil, fmt.Errorf("Missing URL")
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching resources: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %w", err)
	}
	return body, nil
}
