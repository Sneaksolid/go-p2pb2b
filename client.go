package p2pb2b

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIClient struct {
	apiKey    string
	apiSecret string
}

const API_URL = "https://api.p2pb2b.io"
const API_KEY_HEADER = "X-TXC-APIKEY"
const API_PAYLOAD_HEADER = "X-TXC-PAYLOAD"
const API_SIGNATURE_HEADER = "X-TXC-SIGNATURE"

func NewClient(apikey, apiSecret string) *APIClient {
	return &APIClient{
		apiKey:    apikey,
		apiSecret: apiSecret,
	}
}

func (a *APIClient) request(endpoint string, request APIRequest) ([]byte, error) {
	request.SetRequest(endpoint)
	request.SetNonce(time.Now().Unix())

	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(b)
	url := fmt.Sprintf("%v%v", API_URL, endpoint)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	payload := base64.StdEncoding.EncodeToString(b)
	h := hmac.New(sha256.New, []byte(a.apiSecret))
	h.Write([]byte(payload))
	signature := string(h.Sum(nil))

	req.Header.Add(API_KEY_HEADER, a.apiKey)
	req.Header.Add(API_PAYLOAD_HEADER, payload)
	req.Header.Add(API_SIGNATURE_HEADER, signature)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
