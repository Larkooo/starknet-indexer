// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/cartridge-gg/starknet-indexer/ent/event"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/dontpanicdao/caigo/types"
)

// Event is the model entity for the Event schema.
type Event struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// From holds the value of the "from" field.
	From string `json:"from,omitempty"`
	// Keys holds the value of the "keys" field.
	Keys []*types.Felt `json:"keys,omitempty"`
	// Data holds the value of the "data" field.
	Data []*types.Felt `json:"data,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EventQuery when eager-loading is set.
	Edges              EventEdges `json:"edges"`
	transaction_events *string
}

// EventEdges holds the relations/edges for other nodes in the graph.
type EventEdges struct {
	// Transaction holds the value of the transaction edge.
	Transaction *Transaction `json:"transaction,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]*int
}

// TransactionOrErr returns the Transaction value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EventEdges) TransactionOrErr() (*Transaction, error) {
	if e.loadedTypes[0] {
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
func (*Event) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case event.FieldKeys, event.FieldData:
			values[i] = new([]byte)
		case event.FieldID, event.FieldFrom:
			values[i] = new(sql.NullString)
		case event.ForeignKeys[0]: // transaction_events
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Event", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Event fields.
func (e *Event) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case event.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case event.FieldFrom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				e.From = value.String
			}
		case event.FieldKeys:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field keys", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Keys); err != nil {
					return fmt.Errorf("unmarshal field keys: %w", err)
				}
			}
		case event.FieldData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field data", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Data); err != nil {
					return fmt.Errorf("unmarshal field data: %w", err)
				}
			}
		case event.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_events", values[i])
			} else if value.Valid {
				e.transaction_events = new(string)
				*e.transaction_events = value.String
			}
		}
	}
	return nil
}

// QueryTransaction queries the "transaction" edge of the Event entity.
func (e *Event) QueryTransaction() *TransactionQuery {
	return (&EventClient{config: e.config}).QueryTransaction(e)
}

// Update returns a builder for updating this Event.
// Note that you need to call Event.Unwrap() before calling this method if this Event
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Event) Update() *EventUpdateOne {
	return (&EventClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Event entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Event) Unwrap() *Event {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Event is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Event) String() string {
	var builder strings.Builder
	builder.WriteString("Event(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", from=")
	builder.WriteString(e.From)
	builder.WriteString(", keys=")
	builder.WriteString(fmt.Sprintf("%v", e.Keys))
	builder.WriteString(", data=")
	builder.WriteString(fmt.Sprintf("%v", e.Data))
	builder.WriteByte(')')
	return builder.String()
}

// Events is a parsable slice of Event.
type Events []*Event

func (e Events) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
