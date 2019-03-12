package main

import (
	"fmt"
	"github.com/yuexcom/go-xrp"
	"log"
)

func main() {

	client, err := xrp.Dial(xrp.TestNetURL, true)

	if err != nil {
		log.Fatal("dial: ", err)
	}

	cls := &xrp.CommandLedgerStream{
		Streams: []string{"ledger", "transactions"},
	}

	ledgers, err := client.GetLedgers(cls)

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
