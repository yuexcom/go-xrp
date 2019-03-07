package main

import (
	"fmt"
	"log"

	"github.com/yuexcom/go-xrp"
)

func main() {

	client, err := xrp.Dial("s2.ripple.com:443")

	if err != nil {
		log.Fatal("dial: ", err)
	}

	ledgers, err := client.GetLedgers()

	if err != nil {
		log.Fatal("get ledgers: ", err)
	}

	done := make(chan bool)

	go func() {
		for {
			ledger := <-ledgers

			fmt.Printf("-> Ledger: %s Total Tx: %d\n", ledger.LedgerHash, ledger.TxnCount)

			for i, tx := range ledger.Transactions {
				fmt.Printf("--> %d - Transaction: %s\n", i+1, tx.Hash)
				fmt.Printf("---> Addr: %s - Destination Tag: %d\n", tx.Account, tx.DestinationTag)
			}

		}
	}()

	<-done

}
