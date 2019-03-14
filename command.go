package xrp

import "encoding/json"

type TxOptions struct {
	TransactionType string `json:"TransactionType"`
	Account         string `json:"Account"`
	Destination     string `json:"Destination"`
	DestinationTag  uint32 `json:"DestinationTag"`
	Amount          string `json:"Amount"`
	Secret          string `json:"-"`
	Offline         bool   `json:"-"`
	FeeMultMax      int32  `json:"-"`
}

type CommandTX struct {
	Seed       string     `json:"seed,omitempty"`
	SeedHex    string     `json:"seed_hex,omitempty"`
	Passphrase string     `json:"passphrase,omitempty"`
	TxJSON     *TxOptions `json:"tx_json"`
	Secret     string     `json:"secret,omitempty"`
	Offline    bool       `json:"offline,omitempty"`
	FeeMultMax int32      `json:"fee_mult_max,omitempty"`
}

type CommandLedgerStream struct {
	Streams      []string `json:"streams,omitempty"`
	LedgerIndex  string   `json:"ledger_index,omitempty"`
	Full         bool     `json:"full,omitempty"`
	Accounts     bool     `json:"accounts,omitempty"`
	Transactions bool     `json:"transactions,omitempty"`
	Expand       bool     `json:"expand,omitempty"`
	OwnerFunds   bool     `json:"owner_funds,omitempty"`
}

type CommandGetTX struct {
	Hash   string `json:"transaction,omitempty"`
	Binary bool   `json:"binary,omitempty"`
}

type CommandLedger struct {
	LedgerHash  string `json:"ledger_hash,omitempty"`
	LedgerIndex string `json:"ledger_index,omitempty"`
	Full        bool   `json:"full,omitempty"`
	Accounts    bool   `json:"accounts,omitempty"`
	Expand      bool   `json:"expand,omitempty"`
	OwnerFunds  bool   `json:"owner_funds,omitempty"`
	Queue       bool   `json:"queue,omitempty"`
	Binary      bool   `json:"binary,omitempty"`
}

//Command ..
type Command struct {
	Command string `json:"command,omitempty"`
	ID      int    `json:"id,omitempty"`

	*CommandTX
	*CommandGetTX
	*CommandLedger
	*CommandLedgerStream
}

func (cmd *Command) toJSON() (value []byte) {
	value, _ = json.Marshal(cmd)
	return
}
