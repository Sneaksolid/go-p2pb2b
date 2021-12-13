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
	PaginatedResponse
}

func (o *OrderBookResponse) Next() (*OrderBookEntry, error, bool) {
	res, err, ok := o.PaginatedResponse.Next()
	if !ok || err != nil {
		return nil, err, ok
	}

	b, _ := json.Marshal(res)
	entry := new(OrderBookEntry)
	err = json.Unmarshal(b, entry)
	return entry, err, ok
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

	if offset == nil && limit == nil {
		resp.Result.paginationFunc = func(offset int, limit int) ([]byte, error) {
			params["offset"] = strconv.Itoa(offset)
			params["limit"] = strconv.Itoa(limit)

			return a.requestPublicParams(BOOK_ENDPOINT, params)
		}
	}

	return resp.Result, nil
}
