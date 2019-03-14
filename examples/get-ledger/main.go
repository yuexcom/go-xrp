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

	cmd := &xrp.CommandLedger{
		LedgerHash: "DF479D2BEA467660BABCC12DF798517ACCAD94BC5AB8B8274D611C3D0F2A5026",
	}

	ledger, err := client.GetLedger(cmd)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(ledger)
}
