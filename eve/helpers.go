package eve

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var hc = http.Client{
	Timeout: time.Duration(1) * time.Second,
}

func unmarshalJSON(f io.Reader) (map[string]interface{}, error) {
	body, err := io.ReadAll(f)

	if err != nil {
		return map[string]interface{}{}, err
	}

	var marshalledJson map[string]interface{}
	if err := json.Unmarshal(body, &marshalledJson); err != nil {
		return map[string]interface{}{}, err
	}

	return marshalledJson, nil
}

func encodeJSON(f any) (bytes.Buffer, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(f)

	if err != nil {
		return buf, err
	}

	return buf, nil
}

func (c *Client) makeRequest(route string, method string, payload any) (*http.Response, map[string]interface{}, error) {
	p, err := encodeJSON(payload)

	if err != nil {
		return &http.Response{}, map[string]interface{}{}, err
	}

	req, err := http.NewRequest(
		method,
		fmt.Sprintf("%s/%s", c.Instance, route),
		&p,
	)

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))

	if err != nil {
		return &http.Response{}, map[string]interface{}{}, err
	}

	resp, err := hc.Do(req)

	if err != nil {
		return &http.Response{}, map[string]interface{}{}, err
	}

	body, err := unmarshalJSON(resp.Body)

	if err != nil {
		return &http.Response{}, map[string]interface{}{}, err
	}

	return resp, body, nil
}
