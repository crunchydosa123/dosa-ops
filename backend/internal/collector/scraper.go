package collector

import (
	"io"
	"net/http"
)

func scrape(url string) ([]byte, error) {

	resp, err := http.Get(url + "/metrics")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
