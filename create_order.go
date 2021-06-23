package p2pb2b

import "encoding/json"

const CREATE_ORDER_ENDPOINT = "/api/v2/order/new"

type createOrderResponse struct {
	Response
	Result *CreateOrderResponse
}

type CreateOrderRequest struct {
	Request
	Market string `json:"market"`
	Side   string `json:"side"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

type CreateOrderResponse struct {
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

func (a *APIClient) CreateOrder(orderRequest *CreateOrderRequest) (*CreateOrderResponse, error) {
	b, err := a.request(CREATE_ORDER_ENDPOINT, orderRequest)
	if err != nil {
		return nil, err
	}

	resp := new(createOrderResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
