package p2pb2b

import (
	"encoding/json"
)

const BALANCES_ENDPOINT = "/api/v2/account/balances"

type balancesResponse struct {
	Response
	Result map[string]*BalanceResponse
}

type BalanceResponse struct {
	Available string `json:"available"`
	Freeze    string `json:"freeze"`
}

func (a *APIClient) Balances() (map[string]*BalanceResponse, error) {
	req := new(Request)
	b, err := a.request(BALANCES_ENDPOINT, req)
	if err != nil {
		return nil, err
	}

	resp := new(balancesResponse)
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
