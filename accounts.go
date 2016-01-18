package lendingclub

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/shopspring/decimal"
)

const (
	accountsResourcePath  = "/accounts/%d"
	summaryEndpoint       = "/summary"
	availableCashEndpoint = "/availablecash"
	addFundsEndpoint      = "/funds/add"
	withdrawFundsEndpoint = "/funds/withdraw"
	pendingFundsEndpoint  = "/funds/pending"
	cancelFundsEndpoint   = "/funds/cancel"
	notesEndpoint         = "/notes"
	portfoliosEndpoint    = "/portfolios"
	ordersEndpoint        = "/orders"
)

type AccountsResource struct {
	client   *Client
	endpoint string
}

func (c *Client) Accounts(investorID int) *AccountsResource {
	return &AccountsResource{
		client:   c,
		endpoint: fmt.Sprintf(c.baseURL+accountsResourcePath, investorID),
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

type Deposit struct {
	FundsPayload
	InvestorID                 int  `json:"investorId"`
	EstimatedFundsTransferDate Time `json:"estimatedFundsTransferDate"`
}

func (ar *AccountsResource) AddFunds(fundTransfer *FundsPayload) (*Deposit, error) {
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

	var deposit Deposit
	if err := ar.client.processResponse(res, &deposit); err != nil {
		return nil, err
	}

	return &deposit, nil
}

type Withdrawal struct {
	Amount                     decimal.Decimal `json:"amount"`
	InvestorID                 int             `json:"investorId"`
	EstimatedFundsTransferDate Time            `json:"estimatedFundsTransferDate"`
}

func (ar *AccountsResource) WithdrawFunds(amount decimal.Decimal) (*Withdrawal, error) {
	withdrawal := struct {
		Amount decimal.Decimal `json:"amount"`
	}{
		Amount: amount,
	}
	payload, err := json.Marshal(withdrawal)
	if err != nil {
		return nil, err
	}

	req, err := ar.client.newRequest("POST", ar.endpoint+withdrawFundsEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var wd Withdrawal
	if err := ar.client.processResponse(res, &wd); err != nil {
		return nil, err
	}

	return &wd, nil
}

type Transfer struct {
	TransferID    int             `json:"transferId"`
	TransferDate  Time            `json:"transferDate"`
	Amount        decimal.Decimal `json:"amount"`
	SourceAccount string          `json:"sourceAccount"`
	Status        string          `json:"status"`
	Frequency     string          `json:"frequency"`
	EndDate       Time            `json:"endDate"`
	Operation     string          `json:"operation"`
	Cancellable   bool            `json:"cancellable"`
}

func (ar *AccountsResource) PendingFunds() ([]Transfer, error) {
	req, err := ar.client.newRequest("GET", ar.endpoint+pendingFundsEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var respPayload struct {
		Transfers map[int]Transfer `json:"transfers"`
	}
	if err := ar.client.processResponse(res, &respPayload); err != nil {
		return nil, err
	}

	transfers := make([]Transfer, len(respPayload.Transfers))
	for _, transfer := range respPayload.Transfers {
		transfers = append(transfers, transfer)
	}

	return transfers, nil
}

type CancellationResult struct {
	InvestorID    int            `json:"investorId"`
	Cancellations []Cancellation `json:"cancellationResults"`
}

type Cancellation struct {
	TransferID int    `json:"transferId"`
	Status     string `json:"status"`
	Message    string `json:"message"`
}

func (ar *AccountsResource) CancelFunds(transferIds []int) (*CancellationResult, error) {
	transfers := struct {
		TransferIDs []int `json:"transferIds"`
	}{
		TransferIDs: transferIds,
	}
	payload, err := json.Marshal(transfers)
	if err != nil {
		return nil, err
	}

	req, err := ar.client.newRequest("POST", ar.endpoint+cancelFundsEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var cr CancellationResult
	if err := ar.client.processResponse(res, &cr); err != nil {
		return nil, err
	}

	return &cr, nil
}

type Note struct {
	ID               decimal.Decimal `json:"noteId"`
	Amount           decimal.Decimal `json:"noteAmount"`
	LoanID           decimal.Decimal `json:"loanId"`
	OrderID          decimal.Decimal `json:"orderId"`
	InterestRate     decimal.Decimal `json:"interestRate"`
	LoanStatus       string          `json:"loanStatus"`
	Grade            string          `json:"grade"`
	LoanAmount       decimal.Decimal `json:"loanAmount"`
	LoanLength       int             `json:"loanLength"`
	OrderDate        Time            `json:"orderDate"`
	PaymentsReceived decimal.Decimal `json:"paymentsReceived"`
	// TODO: this may be nullable so Time should be a pointer
	IssueDate      Time `json:"issueDate"`
	LoanStatusDate Time `json:"loanStatusDate"`
}

// TODO: Detailed Notes Owned
func (ar *AccountsResource) Notes() ([]Note, error) {
	req, err := ar.client.newRequest("GET", ar.endpoint+notesEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var myNotes struct {
		Notes []Note `json:"myNotes"`
	}
	if err := ar.client.processResponse(res, &myNotes); err != nil {
		return nil, err
	}

	return myNotes.Notes, nil
}

type Portfolio struct {
	ID          int    `json:"portfolioId,omitempty"`
	Name        string `json:"portfolioName"`
	Description string `json:"portfolioDescription,omitempty"`
}

func (ar *AccountsResource) Portfolios() ([]Portfolio, error) {
	req, err := ar.client.newRequest("GET", ar.endpoint+portfoliosEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var myPortfolios struct {
		Portfolios []Portfolio `json:"myPortfolios"`
	}
	if err := ar.client.processResponse(res, &myPortfolios); err != nil {
		return nil, err
	}

	return myPortfolios.Portfolios, nil
}

func (ar *AccountsResource) CreatePortfolio(name, description string) (*Portfolio, error) {
	payload, err := json.Marshal(Portfolio{Name: name, Description: description})
	if err != nil {
		return nil, err
	}

	req, err := ar.client.newRequest("POST", ar.endpoint+portfoliosEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var portfolio Portfolio
	if err := ar.client.processResponse(res, &portfolio); err != nil {
		return nil, err
	}

	return &portfolio, nil
}

type OrderSubmission struct {
	LoanID      int             `json:"loanId"`
	Amount      decimal.Decimal `json:"requestedAmount"`
	PortfolioID int             `json:"portfolioId,omitempty"`
}

type OrderConfirmation struct {
	LoanID          int             `json:"loanId"`
	RequestedAmount decimal.Decimal `json:"requestedAmount"`
	InvestedAmount  int             `json:"investedAmount"`
	ExecutionStatus string          `json:"executionStatus"`
}

type OrderInstruct struct {
	ID                 int                 `json:"orderInstructId"`
	OrderConfirmations []OrderConfirmation `json:"orderConfirmations"`
}

func (ar *AccountsResource) SubmitOrder(accountID int, orders []OrderSubmission) (*OrderInstruct, error) {
	orderSubmission := struct {
		Orders    []OrderSubmission `json:"orders"`
		AccountID int               `json:"aid"`
	}{
		Orders:    orders,
		AccountID: accountID,
	}

	payload, err := json.Marshal(orderSubmission)
	if err != nil {
		return nil, err
	}

	req, err := ar.client.newRequest("POST", ar.endpoint+ordersEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	res, err := ar.client.Do(req)
	if err != nil {
		return nil, err
	}

	var orderInstruct OrderInstruct
	if err := ar.client.processResponse(res, &orderInstruct); err != nil {
		return nil, err
	}

	return &orderInstruct, nil
}
