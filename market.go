package p2pb2b

import "encoding/json"

const MARKET_ENDPOINT = "/api/v2/public/market"

type marketResponse struct {
	Response
	Result *Market
}

func (a *APIClient) Market(market string) (*Market, error) {
	b, err := a.requestPublicParams(MARKET_ENDPOINT, map[string]string{
		"market": market,
	})
	if err != nil {
		return nil, err
	}

	resp := new(marketResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
