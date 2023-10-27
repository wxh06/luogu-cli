package luogu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Request[T any](method, url string, payload any) (data T, err error) {
	reqBody := new(bytes.Buffer)
	if payload != nil {
		err = json.NewEncoder(reqBody).Encode(payload)
		if err != nil {
			return
		}
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return
	}
	req.Header.Add("X-Luogu-Type", "content-only")

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
