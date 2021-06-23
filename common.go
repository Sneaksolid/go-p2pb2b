package p2pb2b

var Debug = false

type APIRequest interface {
	SetRequest(request string)
	SetNonce(nonce int64)
}

type Request struct {
	Request string `json:"request"`
	Nonce   int64  `json:"nonce"`
}

type Response struct {
	Success   bool   `json:"success"`
	ErrorCode int64  `json:"errorCode"`
	Message   string `json:"message"`
}

func (r *Request) SetRequest(request string) {
	r.Request = request
}

func (r *Request) SetNonce(nonce int64) {
	r.Nonce = nonce
}
