package client

import (
	"errors"
	"io"
	"net/http"
)

type httpClient struct {
	Url     string
	headers map[string]string
}

func NewHttpClient(url string, headers map[string]string) *httpClient {

	return &httpClient{
		Url:     url,
		headers: headers,
	}
}

func (client *httpClient) Get() (io.ReadCloser, error) {
	c := &http.Client{}
	req, err := http.NewRequest("GET", client.Url, nil)
	if err != nil {
		return nil, errors.New("Get: " + err.Error())
	}

	for key, val := range client.headers {
		req.Header.Add(key, val)
	}

	res, err := c.Do(req)

	if err != nil {
		return nil, errors.New("Get: " + err.Error())
	}
	return res.Body, nil
}

func (client *httpClient) Post(body io.Reader) (io.ReadCloser, error) {
	c := &http.Client{}

	req, err := http.NewRequest("POST", client.Url, body)
	if err != nil {
		return nil, errors.New("Post: " + err.Error())
	}

	for key, val := range client.headers {
		req.Header.Add(key, val)
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, errors.New("Post: " + err.Error())
	}

	return res.Body, nil
}
