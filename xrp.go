package xrp

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	//TestNetURL todo..
	TestNetURL = "s.altnet.rippletest.net:51233"
	//MainNetURL todo..
	MainNetURL = "s2.ripple.com:443"
)

const (
	statusSuccess    = "success"
	statusError      = "error"
	typeLedgerClosed = "ledgerClosed"
	typeResponse     = "response"
	typeTransaction  = "transaction"

	tesSuccess = "tesSUCCESS"
	tecDryPath = "tecPATH_DRY"
)

//Options todo..
type Options struct {
	URL string
}

//Client main node client
type Client struct {
	conn *websocket.Conn
	Options
	Response        chan *Response
	Ledger          chan *Ledger
	tmpLedger       *Ledger
	tmpLeftTxnCount int
}

//GetTransaction returns tx by hash
func (c *Client) GetTransaction(hash string) (ledger *Ledger, err error) {

	return
}

//GetLedger returns ledger by hash
func (c *Client) GetLedger(cmdLedger *CommandLedger) (ledger *Ledger, err error) {

	cmd := Command{
		ID:            1,
		Command:       "ledger",
		CommandLedger: cmdLedger,
	}

	err = c.SendCommand(cmd.toJSON())

	sr := Response{}

	err = c.conn.ReadJSON(&sr)

	err = c.checkErr(&sr)

	if err != nil {
		return
	}

	ledger = sr.Result.Ledger

	return
}

//GetClosedLedger returns get last closed ledger
func (c *Client) GetClosedLedger() (ledger *Ledger, err error) {

	cmd := Command{
		ID:      2,
		Command: "ledger_closed",
	}

	err = c.SendCommand(cmd.toJSON())

	sr := Response{}

	err = c.conn.ReadJSON(&sr)

	err = c.checkErr(&sr)

	if err != nil {
		return
	}

	ledger = sr.Result.Ledger

	return
}

//GetLedgers get validated ledgers from network
func (c *Client) GetLedgers(cls *CommandLedgerStream) (ledger <-chan *Ledger, err error) {

	cmd := Command{
		ID:                  2,
		Command:             "subscribe",
		CommandLedgerStream: cls,
	}

	fmt.Println(string(cmd.toJSON()))

	err = c.SendCommand(cmd.toJSON())

	return c.Ledger, err
}

//CreateTransaction todo..
func (c *Client) SubmitTransaction(tx *TxOptions) (rsp *TxResult, err error) {

	cmd := Command{
		ID:      2,
		Command: "submit",
		CommandTX: &CommandTX{
			TxJSON:     tx,
			Secret:     tx.Secret,
			Offline:    tx.Offline,
			FeeMultMax: tx.FeeMultMax,
		},
	}

	err = c.SendCommand(cmd.toJSON())

	if err != nil {
		return
	}

	sr := Response{}

	err = c.conn.ReadJSON(&sr)

	if err != nil {
		return
	}

	err = c.checkErr(&sr)

	if err != nil {
		return
	}

	rsp = sr.Result.TxResult

	return
}

//Ping ping XRP server
func (c *Client) Ping() (err error) {

	cmd := Command{
		Command: "ping",
		ID:      1,
	}

	err = c.SendCommand(cmd.toJSON())

	if err != nil {
		return
	}

	sr := Response{}

	err = c.conn.ReadJSON(&sr)

	if err != nil {
		return
	}

	err = c.checkErr(&sr)

	if err != nil {
		return
	}

	return
}

//SendCommand send ws command
func (c *Client) SendCommand(cmd []byte) (err error) {
	err = c.conn.WriteMessage(websocket.TextMessage, cmd)

	if err != nil {
		return
	}

	return
}

//Dial Dial WSS protocol
func Dial(host string, tls bool) (c Client, err error) {

	c = Client{
		Ledger:   make(chan *Ledger),
		Response: make(chan *Response),
	}

	scheme := "ws"

	if tls {
		scheme = "wss"
	}

	u := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   "/",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("go-xrp dial: ", err)
		return
	}

	c.conn = conn

	err = c.Ping()

	if err != nil {
		log.Fatal("go-xrp ping: ", err)
		return
	}

	go c.handleMessage()

	return

}

func (c *Client) handleMessage() {

	defer func() {

		//close ws
		c.conn.Close()

		// close channels
		close(c.Response)
		close(c.Ledger)
	}()

	for {

		rsp := Response{}
		err := c.conn.ReadJSON(&rsp)

		if err != nil {
			log.Fatal("go-xrp read ws response: ", err)
		}

		switch {
		case rsp.Type == typeLedgerClosed:
			c.tmpLedger = &Ledger{
				FeeBase:          rsp.FeeBase,
				FeeRef:           rsp.FeeRef,
				LedgerHash:       rsp.LedgerHash,
				LedgerIndex:      rsp.LedgerIndex,
				LedgerTime:       rsp.LedgerTime,
				ReserveInc:       rsp.ReserveInc,
				ReserveBase:      rsp.ReserveBase,
				TxnCount:         rsp.TxnCount,
				Type:             rsp.Type,
				ValidatedLedgers: rsp.ValidatedLedgers,
			}

			if rsp.TxnCount <= 0 {
				c.Ledger <- c.tmpLedger
			}

			c.tmpLeftTxnCount = rsp.TxnCount
		case rsp.Type == typeResponse && rsp.ID == 2 && rsp.Status == statusSuccess:
			c.tmpLedger = rsp.Result.Ledger
		case rsp.Type == typeTransaction && rsp.Validated == true:
			if c.tmpLedger.LedgerIndex == rsp.LedgerIndex {
				c.tmpLedger.Transactions = append(c.tmpLedger.Transactions, rsp.Transaction)
				c.tmpLeftTxnCount--

				if c.tmpLeftTxnCount == 0 {
					c.Ledger <- c.tmpLedger
				}
			}

		default:
			c.Response <- &rsp
		}

	}
}

func (c *Client) checkErr(res *Response) (errMsg error) {

	if res.Status == statusError {
		errMsg = fmt.Errorf("[ERR:%s:%d] %s", res.Error.Error, res.Error.ErrorCode, res.ErrorMessage)
	}

	return

}
