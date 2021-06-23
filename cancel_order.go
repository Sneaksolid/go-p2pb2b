package p2pb2b

import "encoding/json"

const CANCEL_ORDER_ENDPOINT = "/api/v2/order/cancel"

type cancelOrderResponse struct {
	Response
	Result *Order
}

type CancelOrderRequest struct {
	Request
	Market  string `json:"market"`
	OrderId int64  `json:"orderId"`
}

func (a *APIClient) CancelOrder(cancelOrderRequest *CancelOrderRequest) (*Order, error) {
	b, err := a.request(CANCEL_ORDER_ENDPOINT, cancelOrderRequest)
	if err != nil {
		return nil, err
	}

	resp := new(cancelOrderResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
