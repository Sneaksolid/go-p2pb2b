package p2pb2b

import "encoding/json"

const MARKETS_ENDPOINT = "/api/v2/public/markets"

type marketsResponse struct {
	Response
	Result []*Market
}

func (a *APIClient) Markets() ([]*Market, error) {
	b, err := a.requestPublic(MARKETS_ENDPOINT)
	if err != nil {
		return nil, err
	}

	resp := new(marketsResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
