package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Tonkpils/lendingclub"
)

func main() {
	apiKey := os.Getenv("LC_KEY")
	accountID := os.Getenv("LC_ACCOUNT_ID")
	c := lendingclub.NewClient(apiKey, nil)
	ar := c.Accounts(accoutnID)
	sum, err := ar.Summary()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", sum)

	time.Sleep(1 * time.Second)

	ac, err := ar.AvailableCash()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", ac)

	// fp := &lendingclub.FundsPayload{
	// 	Amount:            decimal.New(100, 1),
	// 	TransferFrequency: "LOAD_NOW",
	// }

	// time.Sleep(1 * time.Second)
	// fr, err := ar.AddFunds(fp)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", fr)

	// time.Sleep(1 * time.Second)
	// wd, err := ar.WithdrawFunds(decimal.New(100, 0))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", wd)

	ts, err := ar.PendingFunds()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", ts)

	// cr, err := ar.CancelFunds([]int{})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%+v\n", cr)

	notes, err := ar.Notes()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", notes)

	portfolios, err := ar.Portfolios()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", portfolios)

	loans, err := c.Loans().Listed()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", loans)
}
