package p2pb2b

import (
	"encoding/json"
	"strconv"
)

const BOOK_ENDPOINT = "/api/v2/public/book"

type bookResponse struct {
	Response
	Result *OrderBookResponse
}

type OrderBookResponse struct {
	Offset int               `json:"offset"`
	Limit  int               `json:"limit"`
	Total  int               `json:"total"`
	Orders []*OrderBookEntry `json:"orders"`
}

func (a *APIClient) Book(market string, side string, offset *int, limit *int) (*OrderBookResponse, error) {
	params := make(map[string]string)
	params["market"] = market
	params["side"] = side
	if offset != nil {
		params["offset"] = strconv.Itoa(*offset)
	}
	if limit != nil {
		params["limit"] = strconv.Itoa(*limit)
	}

	b, err := a.requestPublicParams(BOOK_ENDPOINT, params)
	if err != nil {
		return nil, err
	}

	resp := new(bookResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
