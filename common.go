package p2pb2b

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
)

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

type PaginationFunc func(offset int, limit int) ([]byte, error)

type dummyPagnitationResponse struct {
	Response
	Result *PaginatedResponse
}

type PaginatedResponse struct {
	paginationFunc PaginationFunc `json:"-"`
	current        int            `json:"-"`
	refetch        bool           `json:"-"`

	Offset  int           `json:"offset"`
	Limit   int           `json:"limit"`
	Records []interface{} `json:"records"`
}

func (p *PaginatedResponse) Next() (interface{}, error, bool) {
	if p.current >= p.Limit && p.paginationFunc != nil {
		p.refetch = true
		// refetch
		b, err := p.paginationFunc(p.Offset+p.Limit, p.Limit)
		if err != nil {
			return nil, err, false
		}

		dummy := dummyPagnitationResponse{
			Result: p,
		}
		err = json.Unmarshal(b, &dummy)
		if err != nil {
			return nil, err, false
		}

		p.current = 0
	}

	if p.current >= len(p.Records) {
		return nil, nil, false
	}

	res := p.Records[p.current]
	p.current += 1

	return res, nil, true
}

func (p *PaginatedResponse) UnmarshalJSON(data []byte) error {
	var tmp interface{}
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}

	tmpMap, ok := tmp.(map[string]interface{})
	if ok {
		offset, ok := tmpMap["offset"].(float64)
		if !ok {
			return errors.New("response has no field offset")
		}

		limit, ok := tmpMap["limit"].(float64)
		if !ok {
			return errors.New("response has no field offset")
		}

		records, ok := tmpMap["records"].([]interface{})
		if !ok {
			return errors.New("response has no field records")
		}

		p.Offset = int(offset)
		p.Limit = int(limit)
		p.Records = records

		return nil
	}

	tmpArray, ok := tmp.([]interface{})
	if !ok {
		return errors.New("response is neither a struct nor an array")
	}

	if p.Offset == 0 && p.Limit == 0 {
		// init pagination params. should only happen on first unmarshal
		if len(tmpArray) == 0 {
			p.Limit = 1
		} else {
			p.Limit = len(tmpArray)
		}
	} else if p.refetch {
		p.Offset = p.Offset + p.Limit
	}

	p.Records = tmpArray
	return nil
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

type OrderDeal struct {
	Id          int64   `json:"id"`
	DealOrderId int64   `json:"dealOrderId"`
	Time        float64 `json:"time"`
	Fee         string  `json:"fee"`
	Price       string  `json:"price"`
	Amount      string  `json:"amount"`
	Role        int     `json:"role"`
	Deal        string  `json:"deal"`
}

type OrderBookEntry struct {
	Id        int64   `json:"id"`
	Left      string  `json:"left"`
	Market    string  `json:"market"`
	Amount    string  `json:"amount"`
	Type      string  `json:"type"`
	Price     string  `json:"price"`
	Timestamp float64 `json:"timestamp"`
	Side      string  `json:"side"`
	TakerFee  string  `json:"takerFee"`
	MakerFee  string  `json:"makerFee"`
	DealStock string  `json:"dealStock"`
	DealMoney string  `json:"dealMoney"`
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

type SimpleOrderBookEntry struct {
	Price  string
	Amount string
}

type SimpleOrderBook struct {
	Asks []*SimpleOrderBookEntry
	Bids []*SimpleOrderBookEntry
}

type orderBook []*OrderBookEntry

func (o orderBook) Len() int {
	return len(o)
}

func (o orderBook) Less(i, j int) bool {
	iNum, err := strconv.ParseFloat(o[i].Price, 64)
	if err != nil {
		return false
	}

	jNum, err := strconv.ParseFloat(o[i].Price, 64)
	if err != nil {
		return false
	}

	return iNum > jNum
}

func (o orderBook) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

// Function only our company needs
func GetSimpleOrderBook(buyBooks []*OrderBookEntry, sellBooks []*OrderBookEntry) *SimpleOrderBook {
	sellCpy := make([]*OrderBookEntry, len(sellBooks))
	buyCpy := make([]*OrderBookEntry, len(buyBooks))
	copy(sellCpy, sellBooks)
	copy(buyCpy, buyBooks)

	sort.Sort(orderBook(buyCpy))
	sort.Sort(sort.Reverse(orderBook(sellCpy)))

	asks := make([]*SimpleOrderBookEntry, len(sellCpy))
	for i, val := range sellCpy {
		asks[i] = &SimpleOrderBookEntry{
			Price:  val.Price,
			Amount: val.Amount,
		}
	}

	bids := make([]*SimpleOrderBookEntry, len(buyCpy))
	for i, val := range buyCpy {
		bids[i] = &SimpleOrderBookEntry{
			Price:  val.Price,
			Amount: val.Amount,
		}
	}

	result := &SimpleOrderBook{
		Asks: asks,
		Bids: bids,
	}

	return result
}
