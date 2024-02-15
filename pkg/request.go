package pkg

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
)

func Request(method, uri string, reqBody []byte) *http.Response {
	req := createRequest(method, uri, reqBody)
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp
}

func createRequest(method, uri string, reqBody []byte) *http.Request {
	req := httptest.NewRequest(method, "http://localhost:8080"+uri, bytes.NewBuffer(reqBody))
	req.RequestURI = ""
	req.Header.Set("Content-Type", "application/json")

	return req
}

func DecodeBody(resp *http.Response, dto interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return err
	}

	return nil
}
