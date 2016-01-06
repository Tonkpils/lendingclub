package lendingclub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
)

const (
	accountsResource      = "/accounts/%d"
	summaryEndpoint       = "/summary"
	availableCashEndpoint = "/availablecash"
	addFundsEndpoint      = "/funds/add"
)

type Time struct {
	time.Time
}

const timeFormat = "2006-01-02T15:04:05.999-0700"

func (lct *Time) UnmarshalJSON(b []byte) error {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}

	t, err := time.Parse(timeFormat, string(b))
	if err != nil {
		return err
	}
	*lct = Time{Time: t}

	return nil
}

func (lct Time) MarshalJSON() ([]byte, error) {
	ts := fmt.Sprintf("%q", lct.Format(timeFormat))
	return []byte(ts), nil
}

type AccountsResource struct {
	client   *Client
	endpoint string
}

func (c *Client) Accounts(investorID int) *AccountsResource {
	return &AccountsResource{
		client:   c,
		endpoint: fmt.Sprintf(lendingClubAPI+accountsResource, investorID),
	}
}

type AvailableCash struct {
	InvestorID    int
	AvailableCash decimal.Decimal
}

func (ar *AccountsResource) AvailableCash() (*AvailableCash, error) {
	req, err := ar.client.newRequest("GET", ar.endpoint+availableCashEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var ac AvailableCash
	if err := ar.client.processResponse(res, &ac); err != nil {
		return nil, err
	}

	return &ac, nil
}

type Summary struct {
	AvailableCash        decimal.Decimal
	InvestorID           int
	AccruedInterest      decimal.Decimal
	OutstandingPrincipal decimal.Decimal
	AccountTotal         decimal.Decimal
	TotalNotes           int
	TotalPortfolios      int
	InFundingBalance     decimal.Decimal
	ReceivedInterest     decimal.Decimal
	ReceivedPrincipal    decimal.Decimal
	ReceivedLateFees     decimal.Decimal
}

func (ar *AccountsResource) Summary() (*Summary, error) {
	req, err := ar.client.newRequest("GET", ar.endpoint+summaryEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var sum Summary
	if err := ar.client.processResponse(res, &sum); err != nil {
		return nil, err
	}

	return &sum, nil
}

type FundsPayload struct {
	Amount            decimal.Decimal `json:"amount"`
	TransferFrequency string          `json:"transferFrequency"`
	StartDate         *Time           `json:"startDate,omitempty"`
	EndDate           *Time           `json:"endDate,omitempty"`
}

type FundsResponse struct {
	FundsPayload
	InvestorID                 int  `json:"investorId"`
	EstimatedFundsTransferDate Time `json:"estimatedFundsTransferDate"`
}

func (ar *AccountsResource) AddFunds(fundTransfer *FundsPayload) (*FundsResponse, error) {
	payload, err := json.Marshal(fundTransfer)
	if err != nil {
		return nil, err
	}

	req, err := ar.client.newRequest("POST", ar.endpoint+addFundsEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var fr FundsResponse
	if err := ar.client.processResponse(res, &fr); err != nil {
		return nil, err
	}

	return &fr, nil
}
