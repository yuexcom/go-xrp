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

	res, err := client.GetClosedLedger()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.LedgerIndex)
	fmt.Println(res.LedgerHash)

}
