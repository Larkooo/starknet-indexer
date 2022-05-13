// Code generated by entc, DO NOT EDIT.

package transaction

const (
	// Label holds the string label denoting the transaction type in the database.
	Label = "transaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContractAddress holds the string denoting the contract_address field in the database.
	FieldContractAddress = "contract_address"
	// FieldEntryPointSelector holds the string denoting the entry_point_selector field in the database.
	FieldEntryPointSelector = "entry_point_selector"
	// FieldTransactionHash holds the string denoting the transaction_hash field in the database.
	FieldTransactionHash = "transaction_hash"
	// FieldCalldata holds the string denoting the calldata field in the database.
	FieldCalldata = "calldata"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldNonce holds the string denoting the nonce field in the database.
	FieldNonce = "nonce"
	// EdgeBlock holds the string denoting the block edge name in mutations.
	EdgeBlock = "block"
	// EdgeContract holds the string denoting the contract edge name in mutations.
	EdgeContract = "contract"
	// EdgeReceipts holds the string denoting the receipts edge name in mutations.
	EdgeReceipts = "receipts"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// Table holds the table name of the transaction in the database.
	Table = "transactions"
	// BlockTable is the table that holds the block relation/edge.
	BlockTable = "transactions"
	// BlockInverseTable is the table name for the Block entity.
	// It exists in this package in order to avoid circular dependency with the "block" package.
	BlockInverseTable = "blocks"
	// BlockColumn is the table column denoting the block relation/edge.
	BlockColumn = "block_transactions"
	// ContractTable is the table that holds the contract relation/edge. The primary key declared below.
	ContractTable = "contract_transactions"
	// ContractInverseTable is the table name for the Contract entity.
	// It exists in this package in order to avoid circular dependency with the "contract" package.
	ContractInverseTable = "contracts"
	// ReceiptsTable is the table that holds the receipts relation/edge.
	ReceiptsTable = "transaction_receipts"
	// ReceiptsInverseTable is the table name for the TransactionReceipt entity.
	// It exists in this package in order to avoid circular dependency with the "transactionreceipt" package.
	ReceiptsInverseTable = "transaction_receipts"
	// ReceiptsColumn is the table column denoting the receipts relation/edge.
	ReceiptsColumn = "transaction_receipts"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "transaction_events"
)

// Columns holds all SQL columns for transaction fields.
var Columns = []string{
	FieldID,
	FieldContractAddress,
	FieldEntryPointSelector,
	FieldTransactionHash,
	FieldCalldata,
	FieldSignature,
	FieldNonce,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "transactions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"block_transactions",
}

var (
	// ContractPrimaryKey and ContractColumn2 are the table columns denoting the
	// primary key for the contract relation (M2M).
	ContractPrimaryKey = []string{"contract_id", "transaction_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
