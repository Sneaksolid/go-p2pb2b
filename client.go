package p2pb2b

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type APIClient struct {
	apiKey    string
	apiSecret string
}

const API_URL = "https://api.p2pb2b.com"
const API_KEY_HEADER = "X-TXC-APIKEY"
const API_PAYLOAD_HEADER = "X-TXC-PAYLOAD"
const API_SIGNATURE_HEADER = "X-TXC-SIGNATURE"
const API_CONTENT_TYPE = "Content-Type"

func NewClient(apikey, apiSecret string) *APIClient {
	return &APIClient{
		apiKey:    apikey,
		apiSecret: apiSecret,
	}
}

func NewPublicClient() *APIClient {
	return &APIClient{}
}

func (a *APIClient) requestPublic(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%v%v", API_URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(API_CONTENT_TYPE, "application/json")
	return a.doRequest(req)
}

func (a *APIClient) requestPublicParams(endpoint string, params map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%v%v", API_URL, endpoint)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	for key, val := range params {
		query.Add(key, val)
	}

	req.URL.RawQuery = query.Encode()
	req.Header.Add(API_CONTENT_TYPE, "application/json")
	return a.doRequest(req)
}

func (a *APIClient) request(endpoint string, request APIRequest) ([]byte, error) {
	if a.apiKey == "" || a.apiSecret == "" {
		return nil, errors.New("missing credentials")
	}

	request.SetRequest(endpoint)
	request.SetNonce(time.Now().UnixNano())

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
	h := hmac.New(sha512.New, []byte(a.apiSecret))
	h.Write([]byte(payload))
	signature := hex.EncodeToString(h.Sum(nil))

	req.Header.Add(API_KEY_HEADER, a.apiKey)
	req.Header.Add(API_PAYLOAD_HEADER, payload)
	req.Header.Add(API_SIGNATURE_HEADER, signature)
	req.Header.Add(API_CONTENT_TYPE, "application/json")

	return a.doRequest(req)
}

func (a *APIClient) doRequest(req *http.Request) ([]byte, error) {
	var debugReqBytes []byte
	var err error

	if Debug && req.Body != nil {
		debugReqBytes, err = ioutil.ReadAll(req.Body)
		if err != nil {
			return nil, err
		}
		req.Body.Close()

		newBody := bytes.NewReader(debugReqBytes)
		req.Body = io.NopCloser(newBody)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	r := new(Response)
	err = json.Unmarshal(respBytes, &r)
	if err != nil {
		return nil, err
	}

	if Debug {
		log.Printf("P2PB2B Request: (%v), Response: (%v)", string(debugReqBytes), string(respBytes))
	}

	if !r.Success {
		if val, ok := r.ErrorCode.(int64); ok {
			return nil, fmt.Errorf("API ERROR %d: %v", val, r.Message)
		} else {

			return nil, fmt.Errorf("API ERROR %v: %v", r.ErrorCode, r.Message)
		}
	}

	return respBytes, nil
}
