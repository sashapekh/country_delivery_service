package apiparser

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Request struct {
	Url     string
	Method  string
	Headers map[string]string
	Body    map[string]interface{}
	Params  map[string]string
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (p *Client) MakeRequest(req Request) ([]byte, error) {

	switch req.Method {
	case "GET":
		return p.get(req)
	case "POST":
		return p.post(req)
	default:
		return nil, errors.New("method not supported")
	}
}

func (p *Client) post(req Request) ([]byte, error) {
	preparedBody, err := mapToJSONreader(req.Body)

	if err != nil {
		return nil, err
	}

	resp, err := http.Post(req.Url, "application/json", preparedBody)

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

func (p *Client) get(req Request) ([]byte, error) {
	query := ""

	for key, value := range req.Params {
		query += key + "=" + value + "&"
	}

	requestUrl := req.Url + "?" + query

	resp, err := http.Get(requestUrl)

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

func mapToJSONreader(m map[string]interface{}) (io.Reader, error) {
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}
