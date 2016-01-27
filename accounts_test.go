package lendingclub

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	TestAccountID = 1234
)

func TestAvailableCash(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		availableCashAPI := fmt.Sprintf("/accounts/%d/availablecash", TestAccountID)
		assert.Equal(t, availableCashAPI, req.RequestURI)

		err := respondWithFixture(w, "available_cash.json")
		require.NoError(t, err)
	}))
	defer ts.Close()

	ar := newClient(ts.URL, "Token", nil).Accounts(TestAccountID)
	ac, err := ar.AvailableCash()
	require.NoError(t, err)

	assert.Equal(t, 12345, ac.InvestorID)
	assert.Equal(t, decimal.NewFromFloat(100.76), ac.AvailableCash)
}

func TestSummary(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		summaryAPI := fmt.Sprintf("/accounts/%d/summary", TestAccountID)
		assert.Equal(t, summaryAPI, req.RequestURI)

		err := respondWithFixture(w, "summary.json")
		require.NoError(t, err)
	}))
	defer ts.Close()

	ar := newClient(ts.URL, "Token", nil).Accounts(TestAccountID)
	summary, err := ar.Summary()
	require.NoError(t, err)

	assert.Equal(t, 1788402, summary.InvestorID)
	assert.Equal(t, decimal.NewFromFloat(50.77), summary.AvailableCash)
	assert.Equal(t, decimal.NewFromFloat(100.15), summary.AccountTotal)
	assert.Equal(t, decimal.NewFromFloat(0.26), summary.AccruedInterest)
	assert.Equal(t, decimal.NewFromFloat(0), summary.InFundingBalance)
	assert.Equal(t, decimal.NewFromFloat(0.16), summary.ReceivedInterest)
	assert.Equal(t, decimal.NewFromFloat(0.62), summary.ReceivedPrincipal)
	assert.Equal(t, decimal.NewFromFloat(0), summary.ReceivedLateFees)
	assert.Equal(t, decimal.NewFromFloat(49.38), summary.OutstandingPrincipal)
	assert.Equal(t, 2, summary.TotalNotes)
	assert.Equal(t, 3, summary.TotalPortfolios)
}
