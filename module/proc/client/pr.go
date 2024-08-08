package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PrClient struct {
	host       string
	httpClient *http.Client
}

const (
	getPrURL  = "getPr"
	postPrURL = "postPr"
)

func NewPrClient(host string, httpClient *http.Client) PrClient {
	return PrClient{
		host:       host,
		httpClient: httpClient,
	}
}

func (c *PrClient) GetPr() (resp *GetPrResponse, err error) {
	res, err := do(c.host, *c.httpClient, http.MethodGet, getPrURL, nil, nil)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	response := GetPrResponse{}

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, err
}

func (c *PrClient) PostPr(reqBody PostPrRequest) (resp *PostPrResponse, err error) {
	// payload := fmt.Sprintf(`{"name":"%s"}"`, reqBody.Name)
	// res, err := do(c.host, *c.httpClient, http.MethodPost, postPrURL, nil, payload)

	var body map[string]string
	i, _ := json.Marshal(reqBody)
	json.Unmarshal(i, &body)

	res, err := do(c.host, *c.httpClient, http.MethodPost, postPrURL, nil, body)

	if err != nil {
		fmt.Println(err)
	}

	defer res.Body.Close()

	response := PostPrResponse{}

	err = json.NewDecoder(res.Body).Decode(&response)

	if err != nil {
		return nil, err
	}

	return &response, err
}

func StructToJson(val interface{}) string {
	b, err := json.Marshal(val)
	if err != nil {
		return err.Error()
	}

	return string(b)
}
