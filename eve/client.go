package eve

import (
	"encoding/json"
	"io"
)

// Authenticate uses an email and password to login via Eve and returns a token on success or error on failure
func Authenticate(email string, password string) (string, error) {
	payload, err := encodeJSON(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return "", err
	}

	r, err := hc.Post("http://10.10.9.4:3000/login", "application/json; charset=utf-8", &payload)

	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		return "", err
	}

	var marshalledJson map[string]string
	if err := json.Unmarshal(body, &marshalledJson); err != nil {
		return "", err
	}

	return marshalledJson["token"], nil
}

// Equivalent to the Authenticate function but panics on errors
func MustAuthenticate(email string, password string) string {
	str, err := Authenticate(email, password)

	if err != nil {
		panic(err)
	}

	return str
}

type Client struct {
	Instance string
	Token    string
}
