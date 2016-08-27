package common

import (
	"io"
	"io/ioutil"
	"net/http"
)

func MakeRequestAndGetBytes(url, method string, reader io.Reader) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-length", "0")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
