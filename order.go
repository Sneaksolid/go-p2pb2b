package p2pb2b

import "encoding/json"

const ORDER_ENDPOINT = "/api/v2/account/order"

type orderResponse struct {
	Response
	Result *OrderResponse
}

type OrderResponse struct {
	PaginatedResponse
}

func (o *OrderResponse) Next() (*OrderDeal, error, bool) {
	res, err, ok := o.PaginatedResponse.Next()
	if !ok || err != nil {
		return nil, err, ok
	}

	b, _ := json.Marshal(res)
	deal := new(OrderDeal)
	err = json.Unmarshal(b, deal)
	return deal, err, ok
}

type OrderRequest struct {
	Request
	OrderId int64 `json:"orderId"`
	Offset  int   `json:"offset"`
	Limit   int   `json:"limit"`
}

func (a *APIClient) Order(orderRequest *OrderRequest) (*OrderResponse, error) {
	b, err := a.request(ORDER_ENDPOINT, orderRequest)
	if err != nil {
		return nil, err
	}

	resp := new(orderResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	resp.Result.paginationFunc = func(offset int, limit int) ([]byte, error) {
		orderRequest.Offset = offset
		orderRequest.Limit = limit

		return a.request(ORDER_ENDPOINT, orderRequest)
	}

	return resp.Result, nil
}
