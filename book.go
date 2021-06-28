package p2pb2b

import (
	"encoding/json"
	"strconv"
)

const BOOK_ENDPOINT = "/api/v2/public/book"

type bookResponse struct {
	Response
	Result []*OrderBookEntry
}

func (a *APIClient) Book(market string, side string, offset *int, limit *int) ([]*OrderBookEntry, error) {
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
	return nil, err
}
