package lendingclub

import "github.com/shopspring/decimal"

const (
	loansResourcePath   = "/loans"
	listedLoansEndpoint = "/listing"
)

type LoansResource struct {
	client   *Client
	endpoint string
}

func (c *Client) Loans() *LoansResource {
	return &LoansResource{
		client:   c,
		endpoint: c.baseURL + loansResourcePath,
	}
}

type Loans struct {
	AsOfDate Time   `json:"asOfDate"`
	Loans    []Loan `json:"loans"`
}

type Loan struct {
	ID                                       int             `json:"id"`
	MemberID                                 int             `json:"memberId"`
	Term                                     int             `json:"term"`
	InterestRate                             decimal.Decimal `json:"intRate"`
	ExpectedDefaultRate                      decimal.Decimal `json:"expDefaultRate"`
	ServiceFeeRate                           decimal.Decimal `json:"serviceFeeRate"`
	Installment                              decimal.Decimal `json:"installment"`
	Grade                                    string          `json:"grade"`
	SubGrade                                 string          `json:"subGrade"`
	EmploymentLength                         *int            `json:"empLength"`
	HomeOwnership                            string          `json:"homeOwnership"`
	AnnualIncome                             decimal.Decimal `json:"annualInc"`
	IsIncomeVerified                         string          `json:"isIncV"`
	AcceptDate                               Time            `json:"acceptD"`
	ExpireDate                               Time            `json:"expD"`
	ListDate                                 Time            `json:"listD"`
	CreditPullDate                           Time            `json:"creditPullD"`
	ReviewStatusDate                         *Time           `json:"reviewStatusD"`
	ReviewStatus                             string          `json:"reviewStatus"`
	Description                              string          `json:"desc"`
	Purpose                                  string          `json:"purpose"`
	AddressZip                               string          `json:"addrZip"`
	AddressState                             string          `json:"addrState"`
	InvestorCount                            int             `json:"investorCount"`
	InitialListStatusExpireDate              *Time           `json:"ilsExpD"`
	InitialListStatus                        string          `json:"initialListStatus"`
	EmploymentTitle                          string          `json:"empTitle"`
	AccountsNowDelinquent                    int             `json:"accNowDelinq"`
	AccountsOpenPast24Months                 int             `json:"accOpenPast24Mths"`
	BankcardsOpenToBuy                       int             `json:"bcOpenToBuy"`
	PercentBankcardsGreaterThan75            decimal.Decimal `json:"percentBcGt75"`
	BankcardsUtilization                     decimal.Decimal `json:"bcUtil"`
	DebtToIncome                             decimal.Decimal `json:"dti"`
	DelinquenciesIn2Years                    int             `json:"delinq2Yrs"`
	DelinquentAmount                         decimal.Decimal `json:"delinqAmnt"`
	EarliestCreditLine                       *Time           `json:"earliestCrLine"`
	FICORangeLow                             int             `json:"ficoRangeLow"`
	FICORangeHigh                            int             `json:"ficoRangeHigh"`
	InquiriesLast6Months                     int             `json:"incLast6Mths"`
	MonthsSinceLastDelinquency               int             `json:"mthsSinceLastDelinq"`
	MonthsSinceLastRecord                    int             `json:"mthsSinceLastRecord"`
	MonthsSinceRecentInquiry                 int             `json:"mthsSinceRecentInq"`
	MonthsSinceRecentRevolvingDelinquency    int             `json:"mthsSinceRecentRevolDelinq"`
	MonthsSinceRecentBankcard                int             `json:"mthsSinceRecentBc"`
	MortgageAccounts                         int             `json:"mortAcc"`
	OpenAccounts                             int             `json:"openAcc"`
	PublicRecords                            int             `json:"pubRec"`
	TotalBalanceExcludingMortgage            int             `json:"totalBalExMort"`
	RevolvingBalance                         decimal.Decimal `json:"revolBal"`
	RevolvingUtilization                     decimal.Decimal `json:"revolUtil"`
	TotalBankcardLimit                       int             `json:"totalBcLimit"`
	TotalAccounts                            int             `json:"totalAcc"`
	TotalInstallmentHighCreditLimit          int             `json:"totalIHighCreditLimit"`
	RevolvingAccounts                        int             `json:"numRevAccts"`
	MonthsSinceRecentBankcardDelinquency     int             `json:"mthsSinceRecentBcDlq"`
	PublicRecordBankruptcies                 int             `json:"pubRecBankruptcies"`
	AccountsEver120DaysPastDue               int             `json:"numAcctsEver120Ppd"`
	ChargeoffWithin12Months                  int             `json:"chargeoffWithin12Mths"`
	CollectionsIn12MonthsExcludingMedical    int             `json:"collections12MthsExMed"`
	TaxLiens                                 int             `json:"taxLiens"`
	MonthsSinceLastMajorDerogatoryMark       int             `json:"mthsSinceLastMajorDerog"`
	SatisfactoryAccounts                     int             `json:"numSats"`
	AccountsOpenedInPast12Months             int             `json:"numTlOpPast12m"`
	MonthsSinceRecentAccountOpened           int             `json:"moSinRcntTl"`
	TotalHighCreditLimit                     int             `json:"totHiCredLim"`
	TotalCurrentBalance                      int             `json:"totCurBal"`
	AverageCurrentBalance                    int             `json:"avgCurBal"`
	BankcardAccounts                         int             `json:"numBcTl"`
	ActiveBankcardAccounts                   int             `json:"numActvBctl"`
	SatisfactoryBankcardAccounts             int             `json:"numBcSats"`
	PercentTradesNeverDelinquent             int             `json:"pctTlNvrDlq"`
	Accounts90DaysPastDueIn24Months          int             `json:"numTl90gDpd24m"`
	Accounts30DaysPastDueIn2Months           int             `json:"numTl30dpd"`
	Accounts120DaysPastDueIn2Months          int             `json:"numTl120dpd2m"`
	InstallmentAccounts                      int             `json:"numIlTl"`
	MonthsSinceOldestInstallmentAccount      int             `json:"moSinOldIlAcct"`
	ActiveRevolvingTrades                    int             `json:"numActvRevTl"`
	MonthsSinceOldestRevolvingAccount        int             `json:"moSinOldRevTlOp"`
	MonthsSinceRecentRevolvingAccount        int             `json:"moSinRcntRevTlOp"`
	TotalRevolvingHighCreditLimit            int             `json:"totalRevHiLim"`
	RevolvingTradesWithPositiveBalance       int             `json:"numRevTlBalGt0"`
	OpenRevolvingAccounts                    int             `json:"numOpRevTl"`
	TotalCollectionAmounts                   int             `json:"totCollAmt"`
	FundedAmount                             decimal.Decimal `json:"fundedAmount"`
	LoanAmount                               decimal.Decimal `json:"loanAmount"`
	ApplicationType                          string          `json:"applicationType"`
	JointAnnualIncome                        decimal.Decimal `json:"annualIncJoint"`
	JointDebtToIncome                        decimal.Decimal `json:"dtiJoint"`
	IsJointIncomeVerified                    string          `json:"isIncVJoint"`
	OpenTradesInLast6Months                  int             `json:"openAcc6m"`
	ActiveInstallmentsInLast6Months          int             `json:"openIl6m"`
	OpenedInstallmentsInLast12Months         int             `json:"openIl12m"`
	OpenedInstallmentsInLast24Months         int             `json:"openIl24m"`
	MonthsSinceRecentInstallments            int             `json:"mthsSinceRcntIl"`
	TotalInstallmentsBalance                 decimal.Decimal `json:"totalBalIl"`
	InstallmentsUtilization                  decimal.Decimal `json:"iLUtil"`
	OpenedRevolvingTradesInLast12Months      int             `json:"openRv12m"`
	OpenedRevolvingTradesInLast24Months      int             `json:"openRv24m"`
	MaximumCurrentBalanceOnRevolvingAccounts decimal.Decimal `json:"maxBalBc"`
	AllUtilization                           decimal.Decimal `json:"allUtil"`
	PersonalFinancialInquiries               int             `json:"inqFi"`
	CreditUnionTrades                        int             `json:"totalCuTl"`
	CreditInquiriesInLast12Months            int             `json:"inqLast12m"`
}

func (lr *LoansResource) Listed() (*Loans, error) {
	req, err := lr.client.newRequest("GET", lr.endpoint+listedLoansEndpoint, nil)
	if err != nil {
		return nil, err
	}

	res, err := lr.client.Do(req)
	if err != nil {
		return nil, err
	}

	var loans Loans
	if err := lr.client.processResponse(res, &loans); err != nil {
		return nil, err
	}

	return &loans, nil
}
