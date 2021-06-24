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
	Success bool `json:"success"`
	// Needs to be an interface because the api can't decide on a type
	ErrorCode interface{} `json:"errorCode"`
	Message   string      `json:"message"`
}

func (r *Request) SetRequest(request string) {
	r.Request = request
}

func (r *Request) SetNonce(nonce int64) {
	r.Nonce = nonce
}

type Order struct {
	OrderId   int64   `json:"orderId"`
	Market    string  `json:"market"`
	Price     string  `json:"price"`
	Side      string  `json:"side"`
	Type      string  `json:"type"`
	Timestamp float64 `json:"timestamp"`
	DealMoney string  `json:"dealMoney"`
	DealStock string  `json:"dealStock"`
	Amount    string  `json:"amount"`
	TakerFee  string  `json:"takerFee"`
	MakerFee  string  `json:"makerFee"`
	Left      string  `json:"left"`
	DealFee   string  `json:"dealFee"`
}

type Precision struct {
	Money string `json:"money"`
	Stock string `json:"stock"`
	Fee   string `json:"fee"`
}

type Limits struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
	StepSize  string `json:"step_size"`
	MinPrice  string `json:"min_price"`
	MaxPrice  string `json:"max_price"`
	TickSize  string `json:"tick_size"`
	MinTotal  string `json:"min_total"`
}

type Market struct {
	Name      string     `json:"name"`
	Stock     string     `json:"stock"`
	Money     string     `json:"money"`
	Precision *Precision `json:"precision"`
	Limits    *Limits    `json:"limits"`
}
