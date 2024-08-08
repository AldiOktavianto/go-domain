package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetPrResponse struct {
	Status string `json:"status"`
}

type PostPrRequest struct {
	Name string `json:"name"`
}

type PostPrResponse struct {
	Message string `json:"message"`
}

type ProcClient struct {
	PrClient PrClient
}

type ProcClientParams struct {
	URL string
}

func NewProcClient(param ProcClientParams) *ProcClient {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	prClient := NewPrClient(param.URL, client)

	return &ProcClient{
		PrClient: prClient,
	}
}

func do(host string, client http.Client, method, endpoint string, params map[string]string, payload map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", host, endpoint)

	var body bytes.Buffer

	_ = json.NewEncoder(&body).Encode(payload)

	req, err := http.NewRequest(method, baseURL, &body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()

	return client.Do(req)
}
