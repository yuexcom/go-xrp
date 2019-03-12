package main

import (
	"fmt"
	"github.com/yuexcom/go-xrp"
	"log"
)

func main() {

	client, err := xrp.Dial("127.0.0.1:6001", false)

	if err != nil {
		log.Fatal("dial: ", err)
	}

	tx := &xrp.TxOptions{
		TransactionType: "Payment",
		Secret:          "****",
		Account:         "****",
		Destination:     "****",
		DestinationTag:  3544,
		Amount:          "20000",
	}

	res, err := client.SubmitTransaction(tx)

	if err != nil {
		log.Fatal(err)
	}

	if res.EngineResult == "tesSUCCESS" {
		fmt.Println(res.Account)
		fmt.Println(res.Amount)
	}

}
