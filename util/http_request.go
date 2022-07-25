package util

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(urlPath string) ([]byte, error) {
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}
