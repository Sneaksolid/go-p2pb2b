package p2pb2b

import "encoding/json"

const CREATE_ORDER_ENDPOINT = "/api/v2/order/new"

type createOrderResponse struct {
	Response
	Result *Order
}

type CreateOrderRequest struct {
	Request
	Market string `json:"market"`
	Side   string `json:"side"`
	Amount string `json:"amount"`
	Price  string `json:"price"`
}

func (a *APIClient) CreateOrder(orderRequest *CreateOrderRequest) (*Order, error) {
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
