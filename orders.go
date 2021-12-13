package p2pb2b

import "encoding/json"

const ORDERS_ENDPOINT = "/api/v2/orders"

type ordersResponse struct {
	Response
	Result *OrdersResponse
}

type OrdersResponse struct {
	PaginatedResponse
}

func (o *OrdersResponse) Next() (*Order, error, bool) {
	res, err, ok := o.PaginatedResponse.Next()
	if !ok || err != nil {
		return nil, err, ok
	}

	b, _ := json.Marshal(res)
	order := new(Order)
	err = json.Unmarshal(b, order)
	return order, err, ok
}

type OrdersRequest struct {
	Request
	Market string `json:"market"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

func (a *APIClient) Orders(ordersRequest *OrdersRequest) (*OrdersResponse, error) {
	b, err := a.request(ORDERS_ENDPOINT, ordersRequest)
	if err != nil {
		return nil, err
	}

	resp := new(ordersResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	resp.Result.paginationFunc = func(offset int, limit int) ([]byte, error) {
		ordersRequest.Offset = offset
		ordersRequest.Limit = limit

		return a.request(ORDERS_ENDPOINT, ordersRequest)
	}

	return resp.Result, nil
}
