package lendingclub

import (
	"fmt"

	"github.com/shopspring/decimal"
)

const (
	accountsResource      = "/accounts/%d"
	summaryEndpoint       = "/summary"
	availableCashEndpoint = "/availablecash"
)

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
