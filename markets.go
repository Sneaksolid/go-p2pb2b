package p2pb2b

import "encoding/json"

const MARKETS_ENDPOINT = "/api/v2/public/markets"

type marketsResponse struct {
	Response
	Result []*Market
}

type Precision struct {
	Money string `json:"money"`
	Stock string `json:"stock"`
	Fee   string `json:"fee"`
}

type Limits struct {
	MinAmount string `json:"min_amount"`
	MaxAmount string `json:"max_amount"`
	StepSize  string `json:"step_size"`
	MinPrice  string `json:"min_price"`
	MaxPrice  string `json:"max_price"`
	TickSize  string `json:"tick_size"`
	MinTotal  string `json:"min_total"`
}

type Market struct {
	Name      string     `json:"name"`
	Stock     string     `json:"stock"`
	Money     string     `json:"money"`
	Precision *Precision `json:"precision"`
	Limits    *Limits    `json:"limits"`
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
