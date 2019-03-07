package xrp

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

var (
	networkURL = map[string]string{
		"MAIN": "s2.ripple.com:443",
		"TEST": "s.altnet.rippletest.net:51233",
	}
)

const (
	STATUS_SUCCESS     = "success"
	TYPE_LEDGER_CLOSED = "ledgerClosed"
	TYPE_RESPONSE      = "response"
	TYPE_TRANSACTION   = "transaction"

	TES_SUCCESS  = "tesSUCCESS"
	TEC_PATH_DRY = "tecPATH_DRY"
)

// NETWORK

type Info struct {
	BuildVersion string `json:"build_version"`
	HostID       string
	Peers        int
}

type Result struct {
	*Ledger
	*Info
}

type Response struct {
	ID     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Type   string `json:"type,omitempty"`

	Result *Result `json:"result,omitempty"`

	EngineResult        string `json:"engine_result"`
	EngineResultCode    int    `json:"engine_result_code"`
	LedgerCurrentIndex  uint64 `json:"ledger_current_index"`
	EngineResultMessage string `json:"engine_result_message"`
	LedgerHash          string `json:"ledger_hash"`
	LedgerIndex         uint64 `json:"ledger_index"`

	Transaction Transaction `json:"transaction"`

	Validated bool `json:"validated"`

	Ledger
}

//Command
type Command struct {
	Command string   `json:"command,omitempty"`
	ID      int      `json:"id,omitempty"`
	Streams []string `json:"streams,omitempty"`

	LedgerIndex  string `json:"ledger_index,omitempty"`
	Full         bool   `json:"full,omitempty"`
	Accounts     bool   `json:"accounts,omitempty"`
	Transactions bool   `json:"transactions,omitempty"`
	Expand       bool   `json:"expand,omitempty"`
	OwnerFunds   bool   `json:"owner_funds,omitempty"`
}

func (cmd *Command) toJSON() (value []byte) {
	value, _ = json.Marshal(cmd)
	return
}

// OBJS

type TransactionMeta struct {
	AffectedNodes []struct {
		CreatedNode struct {
			LedgerEntryType string `json:"LedgerEntryType"`
			LedgerIndex     string `json:"LedgerIndex"`
			NewFields       struct {
				Account       string `json:"Account"`
				BookDirectory string `json:"BookDirectory"`
				Expiration    int64  `json:"Expiration"`
				Flags         int32  `json:"Flags"`
				Sequence      int64  `json:"Sequence"`
				TakerGets     struct {
					Currency string `json:"currency"`
					Issuer   string `json:"issuer"`
					Value    string `json:"value"`
				} `json:"TakerGets"`
				TakerPays struct {
					Currency string `json:"currency"`
					Issuer   string `json:"issuer"`
					Value    string `json:"value"`
				} `json:"TakerPays"`
			} `json:"NewFields"`
		} `json:"CreatedNode,omitempty"`
		ModifiedNode struct {
			FinalFields struct {
				Flags         int    `json:"Flags"`
				IndexNext     string `json:"IndexNext"`
				IndexPrevious string `json:"IndexPrevious"`
				Owner         string `json:"Owner"`
				RootIndex     string `json:"RootIndex"`
			} `json:"FinalFields"`
			LedgerEntryType string `json:"LedgerEntryType"`
			LedgerIndex     string `json:"LedgerIndex"`
		} `json:"ModifiedNode,omitempty"`
		DeletedNode struct {
			FinalFields struct {
				Account           string `json:"Account"`
				BookDirectory     string `json:"BookDirectory"`
				BookNode          string `json:"BookNode"`
				Expiration        int64  `json:"Expiration"`
				Flags             int32  `json:"Flags"`
				OwnerNode         string `json:"OwnerNode"`
				PreviousTxnID     string `json:"PreviousTxnID"`
				PreviousTxnLgrSeq int64  `json:"PreviousTxnLgrSeq"`
				Sequence          int32  `json:"Sequence"`
				TakerGets         struct {
					Currency string `json:"currency"`
					Issuer   string `json:"issuer"`
					Value    string `json:"value"`
				} `json:"TakerGets"`
				TakerPays struct {
					Currency string `json:"currency"`
					Issuer   string `json:"issuer"`
					Value    string `json:"value"`
				} `json:"TakerPays"`
			} `json:"FinalFields"`
			LedgerEntryType string `json:"LedgerEntryType"`
			LedgerIndex     string `json:"LedgerIndex"`
		} `json:"DeletedNode,omitempty"`
	} `json:"AffectedNodes"`
	TransactionIndex  int64  `json:"TransactionIndex"`
	TransactionResult string `json:"TransactionResult"`
}

//Transaction
type Transaction struct {
	Account            string `json:"Account"`
	Expiration         int64  `json:"Expiration"`
	Fee                string `json:"Fee"`
	Flags              int64  `json:"Flags"`
	LastLedgerSequence int64  `json:"LastLedgerSequence"`
	OfferSequence      int64  `json:"OfferSequence"`
	Sequence           int64  `json:"Sequence"`
	SigningPubKey      string `json:"SigningPubKey"`
	DestinationTag     uint32 `json:"DestinationTag"`
	TransactionType    string `json:"TransactionType"`
	TxnSignature       string `json:"TxnSignature"`
	Date               int64  `json:"date"`
	Hash               string `json:"hash"`
	OwnerFunds         string `json:"owner_funds"`
}

//Ledger
type Ledger struct {
	FeeBase          int    `json:"fee_base,omitempty"`
	FeeRef           int    `json:"fee_ref,omitempty"`
	LedgerHash       string `json:"ledger_hash,omitempty"`
	LedgerIndex      uint64 `json:"ledger_index,omitempty"`
	LedgerTime       int    `json:"ledger_time,omitempty"`
	ReserveBase      int    `json:"reserve_base,omitempty"`
	ReserveInc       int    `json:"reserve_inc,omitempty"`
	TxnCount         int    `json:"txn_count,omitempty"`
	Type             string `json:"type,omitempty"`
	ValidatedLedgers string `json:"validated_ledgers,omitempty"`
	Transactions     []Transaction
}

//OPTS

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

//GetLedgers get validated ledgers from network
func (c *Client) GetLedgers() (ledger <-chan *Ledger, err error) {

	cmd := Command{
		ID:      2,
		Command: "subscribe",
		Streams: []string{"ledger", "transactions"},
	}

	err = c.SendCommand(cmd.toJSON())

	return c.Ledger, err
}

func (c *Client) GetTransactions(hash string) (transactions *[]Transaction, err error) {

	return
}

//Ping ping XRP server
func (c *Client) Ping() (status bool, err error) {

	cmd := Command{
		Command: "ping",
		ID:      1,
	}

	err = c.SendCommand(cmd.toJSON())

	if err != nil {
		return
	}

	rsp := Response{}

	err = c.conn.ReadJSON(&rsp)

	if rsp.Status == STATUS_SUCCESS {
		status = true
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
func Dial(host string) (c Client, err error) {

	c = Client{
		Ledger:   make(chan *Ledger),
		Response: make(chan *Response),
	}

	u := url.URL{
		Scheme: "wss",
		Host:   host,
		Path:   "/",
	}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("go-xrp dial:", err)
		return
	}

	c.conn = conn

	status, err := c.Ping()

	if err != nil {
		log.Fatal("go-xrp ping: ", err)
		return
	}

	if !status {
		log.Fatal("go-xrp ping: ", errPing)
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
		case rsp.Type == TYPE_LEDGER_CLOSED:
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
			c.tmpLeftTxnCount = rsp.TxnCount
		case rsp.Type == TYPE_RESPONSE && rsp.ID == 2 && rsp.Status == STATUS_SUCCESS:
			c.tmpLedger = rsp.Result.Ledger
		case rsp.Type == TYPE_TRANSACTION && rsp.Validated == true:
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
