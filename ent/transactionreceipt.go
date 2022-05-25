// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
	"github.com/dontpanicdao/caigo/types"
)

// TransactionReceipt is the model entity for the TransactionReceipt schema.
type TransactionReceipt struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// Status holds the value of the "status" field.
	Status transactionreceipt.Status `json:"status,omitempty"`
	// StatusData holds the value of the "status_data" field.
	StatusData string `json:"status_data,omitempty"`
	// MessagesSent holds the value of the "messages_sent" field.
	MessagesSent []*types.L1Message `json:"messages_sent,omitempty"`
	// L1OriginMessage holds the value of the "l1_origin_message" field.
	L1OriginMessage *types.L2Message `json:"l1_origin_message,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionReceiptQuery when eager-loading is set.
	Edges                      TransactionReceiptEdges `json:"edges"`
	block_transaction_receipts *string
	transaction_receipt        *string
}

// TransactionReceiptEdges holds the relations/edges for other nodes in the graph.
type TransactionReceiptEdges struct {
	// Block holds the value of the block edge.
	Block *Block `json:"block,omitempty"`
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]*int
}

// BlockOrErr returns the Block value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionReceiptEdges) BlockOrErr() (*Block, error) {
	if e.loadedTypes[0] {
		if e.Block == nil {
			// The edge block was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: block.Label}
		}
		return e.Block, nil
	}
	return nil, &NotLoadedError{edge: "block"}
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionReceiptEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[1] {
		if e.Transaction == nil {
			// The edge transaction was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: transaction.Label}
		}
		return e.Transaction, nil
	}
	return nil, &NotLoadedError{edge: "transaction"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TransactionReceipt) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactionreceipt.FieldMessagesSent, transactionreceipt.FieldL1OriginMessage:
			values[i] = new([]byte)
		case transactionreceipt.FieldID, transactionreceipt.FieldTransactionHash, transactionreceipt.FieldStatus, transactionreceipt.FieldStatusData:
			values[i] = new(sql.NullString)
		case transactionreceipt.ForeignKeys[0]: // block_transaction_receipts
			values[i] = new(sql.NullString)
		case transactionreceipt.ForeignKeys[1]: // transaction_receipt
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TransactionReceipt", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TransactionReceipt fields.
func (tr *TransactionReceipt) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactionreceipt.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				tr.ID = value.String
			}
		case transactionreceipt.FieldTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_hash", values[i])
			} else if value.Valid {
				tr.TransactionHash = value.String
			}
		case transactionreceipt.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				tr.Status = transactionreceipt.Status(value.String)
			}
		case transactionreceipt.FieldStatusData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status_data", values[i])
			} else if value.Valid {
				tr.StatusData = value.String
			}
		case transactionreceipt.FieldMessagesSent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field messages_sent", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tr.MessagesSent); err != nil {
					return fmt.Errorf("unmarshal field messages_sent: %w", err)
				}
			}
		case transactionreceipt.FieldL1OriginMessage:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field l1_origin_message", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &tr.L1OriginMessage); err != nil {
					return fmt.Errorf("unmarshal field l1_origin_message: %w", err)
				}
			}
		case transactionreceipt.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field block_transaction_receipts", values[i])
			} else if value.Valid {
				tr.block_transaction_receipts = new(string)
				*tr.block_transaction_receipts = value.String
			}
		case transactionreceipt.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_receipt", values[i])
			} else if value.Valid {
				tr.transaction_receipt = new(string)
				*tr.transaction_receipt = value.String
			}
		}
	}
	return nil
}

// QueryBlock queries the "block" edge of the TransactionReceipt entity.
func (tr *TransactionReceipt) QueryBlock() *BlockQuery {
	return (&TransactionReceiptClient{config: tr.config}).QueryBlock(tr)
}

// QueryTransaction queries the "transaction" edge of the TransactionReceipt entity.
func (tr *TransactionReceipt) QueryTransaction() *TransactionQuery {
	return (&TransactionReceiptClient{config: tr.config}).QueryTransaction(tr)
}

// Update returns a builder for updating this TransactionReceipt.
// Note that you need to call TransactionReceipt.Unwrap() before calling this method if this TransactionReceipt
// was returned from a transaction, and the transaction was committed or rolled back.
func (tr *TransactionReceipt) Update() *TransactionReceiptUpdateOne {
	return (&TransactionReceiptClient{config: tr.config}).UpdateOne(tr)
}

// Unwrap unwraps the TransactionReceipt entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tr *TransactionReceipt) Unwrap() *TransactionReceipt {
	tx, ok := tr.config.driver.(*txDriver)
	if !ok {
		panic("ent: TransactionReceipt is not a transactional entity")
	}
	tr.config.driver = tx.drv
	return tr
}

// String implements the fmt.Stringer.
func (tr *TransactionReceipt) String() string {
	var builder strings.Builder
	builder.WriteString("TransactionReceipt(")
	builder.WriteString(fmt.Sprintf("id=%v", tr.ID))
	builder.WriteString(", transaction_hash=")
	builder.WriteString(tr.TransactionHash)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", tr.Status))
	builder.WriteString(", status_data=")
	builder.WriteString(tr.StatusData)
	builder.WriteString(", messages_sent=")
	builder.WriteString(fmt.Sprintf("%v", tr.MessagesSent))
	builder.WriteString(", l1_origin_message=")
	builder.WriteString(fmt.Sprintf("%v", tr.L1OriginMessage))
	builder.WriteByte(')')
	return builder.String()
}

// TransactionReceipts is a parsable slice of TransactionReceipt.
type TransactionReceipts []*TransactionReceipt

func (tr TransactionReceipts) config(cfg config) {
	for _i := range tr {
		tr[_i].config = cfg
	}
}
