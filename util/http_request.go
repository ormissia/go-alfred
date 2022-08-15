package util

import (
	"io"
	"net/http"
)

func HttpGet(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(resp.Body)
}
