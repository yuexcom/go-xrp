package xrp

//Info todo..
type Info struct {
	BuildVersion string `json:"build_version"`
	HostID       string
	Peers        int
}

type TxJSON struct {
	Account         string `json:"Account,omitempty"`
	Amount          string `json:"Amount,omitempty"`
	Destination     string `json:"Destination,omitempty"`
	DestinationTag  int    `json:"DestinationTag,omitempty"`
	Fee             string `json:"Fee,omitempty"`
	Flags           int64  `json:"Flags,omitempty"`
	Sequence        int    `json:"Sequence,omitempty"`
	SigningPubKey   string `json:"SigningPubKey,omitempty"`
	TransactionType string `json:"TransactionType,omitempty"`
	TxnSignature    string `json:"TxnSignature,omitempty"`
	Hash            string `json:"hash,omitempty"`
}

type TxResult struct {
	EngineResult        string `json:"engine_result,omitempty"`
	EngineResultCode    int    `json:"engine_result_code,omitempty"`
	EngineResultMessage string `json:"engine_result_message,omitempty"`

	TxBlob  string `json:"tx_blob,omitempty"`
	*TxJSON `json:"tx_json,omitempty"`
}

//Result todo..
type Result struct {
	*Ledger
	*Info
	*TxResult
}

//Request todo..
type Request struct {
	Account     string `json:"account,omitempty"`
	Command     string `json:"command,omitempty"`
	ID          int    `json:"id,omitempty"`
	LedgerIndex string `json:"ledger_index,omitempty"`
	Strict      bool   `json:"strict,omitempty"`
}

type Error struct {
	Error        string `json:"error,omitempty"`
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
}

//Response todo..
type Response struct {
	ID     int    `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
	Type   string `json:"type,omitempty"`

	*Error
	*Request `json:"request,omitempty"`
	*Result  `json:"result,omitempty"`

	// pure tx subs
	EngineResult        string `json:"engine_result,omitempty"`
	EngineResultCode    int    `json:"engine_result_code,omitempty"`
	LedgerCurrentIndex  uint64 `json:"ledger_current_index,omitempty"`
	EngineResultMessage string `json:"engine_result_message,omitempty"`
	LedgerHash          string `json:"ledger_hash,omitempty"`
	LedgerIndex         uint64 `json:"ledger_index,omitempty"`

	Transaction Transaction `json:"transaction,omitempty"`
	Validated   bool        `json:"validated,omitempty"`
	Ledger
}

//Transaction todo..
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

//Ledger todo..
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
