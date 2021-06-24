package p2pb2b

import "encoding/json"

const TICKERS_ENDPOINT = "/api/v2/public/tickers"

type Ticker struct {
	Bid    string `json:"bid"`
	Ask    string `json:"ask"`
	Low    string `json:"low"`
	High   string `json:"high"`
	Last   string `json:"last"`
	Vol    string `json:"vol"`
	Change string `json:"change"`
}

type TickerResponse struct {
	At     int64   `json:"at"`
	Ticker *Ticker `json:"ticker"`
}

type tickersResponse struct {
	Response
	Result map[string]*TickerResponse
}

func (a *APIClient) Tickers() (map[string]*TickerResponse, error) {
	b, err := a.requestPublic(TICKERS_ENDPOINT)
	if err != nil {
		return nil, err
	}

	resp := new(tickersResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
