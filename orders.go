package p2pb2b

import "encoding/json"

const ORDERS_ENDPOINT = "/api/v2/orders"

type ordersResponse struct {
	Response
	Result []*Order
}

type OrdersRequest struct {
	Request
	Market string `json:"market"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (a *APIClient) Orders(ordersRequest *OrdersRequest) ([]*Order, error) {
	b, err := a.request(ORDERS_ENDPOINT, ordersRequest)
	if err != nil {
		return nil, err
	}

	resp := new(ordersResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
