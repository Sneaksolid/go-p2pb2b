package p2pb2b

import "encoding/json"

const ORDER_ENDPOINT = "/api/v2/account/order"

type orderResponse struct {
	Response
	Result *OrderResponse
}

type OrderResponse struct {
	Offset  int          `json:"offset"`
	Limit   int          `json:"limit"`
	Records []*OrderDeal `json:"records"`
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

	return resp.Result, nil
}
