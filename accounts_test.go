package lendingclub

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shopspring/decimal"
)

var (
	TestAccountID = 1234
)

func TestAvailableCash(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		availableCashAPI := fmt.Sprintf("/accounts/%d/availablecash", TestAccountID)
		if availableCashAPI != req.RequestURI {
			t.Errorf("Expected available cash API to be %s, but got %s",
				availableCashAPI, req.RequestURI)
		}

		if err := respondWithFixture(w, "available_cash.json"); err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()

	ar := newClient(ts.URL, "Token", nil).Accounts(TestAccountID)
	ac, err := ar.AvailableCash()
	if err != nil {
		t.Fatal(err)
	}

	if expected := 12345; ac.InvestorID != expected {
		t.Errorf("Expected Investor ID to be %d, but got %d",
			expected, ac.InvestorID)
	}

	if expected := decimal.NewFromFloat(100.76); !ac.AvailableCash.Equals(expected) {
		t.Errorf("Expected Available Cash to equal %s, but got %s",
			expected, ac.AvailableCash)
	}
}
