package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, []byte{}
	}
	req.Header.Set("x-into-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJfbmFtZSI6InRlc3QifQ.riZtv7y-kexYAb7mXp6cpf9G-Flb1rb-2POtNQXXe8E")
	resp, err := client.Do(req)
	if err != nil {
		return err, []byte{}
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, []byte{}
	}

	return nil, body
}

func GetWithHeader(url string, headers map[string]string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err, []byte{}
	}
	for key, header := range headers {
		req.Header.Set(key, header)
	}
	resp, err := client.Do(req)
	if err != nil {
		return err, []byte{}
	}
	defer resp.Body.Close()
	fmt.Println(1)
	body, err := io.ReadAll(resp.Body)
	fmt.Println(2)
	if err != nil {
		return err, []byte{}
	}
	return nil, body
}

func HttpPost(url string) (error, []byte) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err, []byte{}
	}
	req.Header.Set("x-into-token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxMjMsInVzZXJfbmFtZSI6InRlc3QifQ.riZtv7y-kexYAb7mXp6cpf9G-Flb1rb-2POtNQXXe8E")
	resp, err := client.Do(req)
	if err != nil {
		return err, []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, []byte{}
	}

	return nil, body
}
