package luogu

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

	clientId := make([]byte, 20)
	if _, err = rand.Read(clientId); err != nil {
		return
	}
	req.AddCookie(&http.Cookie{Name: "__client_id", Value: fmt.Sprintf("%040s", hex.EncodeToString(clientId))})

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var errData ErrorResponse
	if err = json.Unmarshal(body, &errData); err != nil {
		return
	}
	if errData.Status >= 400 {
		return *new(T), errors.New(errData.ErrorMessage)
	}
	err = json.Unmarshal(body, &data)
	return
}
