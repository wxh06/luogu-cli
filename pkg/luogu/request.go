package luogu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Get[T any](url string) (T, error) {
	return Request[T]("GET", url, nil)
}

func Post[T any](url string, payload any) (data T, err error) {
	reqBody := new(bytes.Buffer)
	err = json.NewEncoder(reqBody).Encode(payload)
	if err != nil {
		return
	}
	return Request[T]("POST", url, reqBody)
}

func Request[T any](method, url string, reqBody io.Reader) (data T, err error) {
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return
	}
	req.Header.Add("X-Luogu-Type", "content-only")
	req.AddCookie(&http.Cookie{Name: "__client_id", Value: "0000000000000000000000000000000000000000"})

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &data)
	return
}
