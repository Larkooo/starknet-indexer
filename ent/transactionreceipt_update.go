// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dontpanicdao/caigo/types"
	"github.com/tarrencev/starknet-indexer/ent/predicate"
	"github.com/tarrencev/starknet-indexer/ent/transaction"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

// TransactionReceiptUpdate is the builder for updating TransactionReceipt entities.
type TransactionReceiptUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionReceiptMutation
}

// Where appends a list predicates to the TransactionReceiptUpdate builder.
func (tru *TransactionReceiptUpdate) Where(ps ...predicate.TransactionReceipt) *TransactionReceiptUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetTransactionHash sets the "transaction_hash" field.
func (tru *TransactionReceiptUpdate) SetTransactionHash(s string) *TransactionReceiptUpdate {
	tru.mutation.SetTransactionHash(s)
	return tru
}

// SetStatus sets the "status" field.
func (tru *TransactionReceiptUpdate) SetStatus(t transactionreceipt.Status) *TransactionReceiptUpdate {
	tru.mutation.SetStatus(t)
	return tru
}

// SetStatusData sets the "status_data" field.
func (tru *TransactionReceiptUpdate) SetStatusData(s string) *TransactionReceiptUpdate {
	tru.mutation.SetStatusData(s)
	return tru
}

// SetMessagesSent sets the "messages_sent" field.
func (tru *TransactionReceiptUpdate) SetMessagesSent(t []*types.L1Message) *TransactionReceiptUpdate {
	tru.mutation.SetMessagesSent(t)
	return tru
}

// SetL1OriginMessage sets the "l1_origin_message" field.
func (tru *TransactionReceiptUpdate) SetL1OriginMessage(t *types.L2Message) *TransactionReceiptUpdate {
	tru.mutation.SetL1OriginMessage(t)
	return tru
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (tru *TransactionReceiptUpdate) SetTransactionID(id string) *TransactionReceiptUpdate {
	tru.mutation.SetTransactionID(id)
	return tru
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (tru *TransactionReceiptUpdate) SetNillableTransactionID(id *string) *TransactionReceiptUpdate {
	if id != nil {
		tru = tru.SetTransactionID(*id)
	}
	return tru
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (tru *TransactionReceiptUpdate) SetTransaction(t *Transaction) *TransactionReceiptUpdate {
	return tru.SetTransactionID(t.ID)
}

// Mutation returns the TransactionReceiptMutation object of the builder.
func (tru *TransactionReceiptUpdate) Mutation() *TransactionReceiptMutation {
	return tru.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (tru *TransactionReceiptUpdate) ClearTransaction() *TransactionReceiptUpdate {
	tru.mutation.ClearTransaction()
	return tru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TransactionReceiptUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tru.hooks) == 0 {
		if err = tru.check(); err != nil {
			return 0, err
		}
		affected, err = tru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionReceiptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tru.check(); err != nil {
				return 0, err
			}
			tru.mutation = mutation
			affected, err = tru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tru.hooks) - 1; i >= 0; i-- {
			if tru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TransactionReceiptUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TransactionReceiptUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TransactionReceiptUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tru *TransactionReceiptUpdate) check() error {
	if v, ok := tru.mutation.Status(); ok {
		if err := transactionreceipt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TransactionReceipt.status": %w`, err)}
		}
	}
	return nil
}

func (tru *TransactionReceiptUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionreceipt.Table,
			Columns: transactionreceipt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactionreceipt.FieldID,
			},
		},
	}
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.TransactionHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldTransactionHash,
		})
	}
	if value, ok := tru.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transactionreceipt.FieldStatus,
		})
	}
	if value, ok := tru.mutation.StatusData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldStatusData,
		})
	}
	if value, ok := tru.mutation.MessagesSent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldMessagesSent,
		})
	}
	if value, ok := tru.mutation.L1OriginMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldL1OriginMessage,
		})
	}
	if tru.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transactionreceipt.TransactionTable,
			Columns: []string{transactionreceipt.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transactionreceipt.TransactionTable,
			Columns: []string{transactionreceipt.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionreceipt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TransactionReceiptUpdateOne is the builder for updating a single TransactionReceipt entity.
type TransactionReceiptUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionReceiptMutation
}

// SetTransactionHash sets the "transaction_hash" field.
func (truo *TransactionReceiptUpdateOne) SetTransactionHash(s string) *TransactionReceiptUpdateOne {
	truo.mutation.SetTransactionHash(s)
	return truo
}

// SetStatus sets the "status" field.
func (truo *TransactionReceiptUpdateOne) SetStatus(t transactionreceipt.Status) *TransactionReceiptUpdateOne {
	truo.mutation.SetStatus(t)
	return truo
}

// SetStatusData sets the "status_data" field.
func (truo *TransactionReceiptUpdateOne) SetStatusData(s string) *TransactionReceiptUpdateOne {
	truo.mutation.SetStatusData(s)
	return truo
}

// SetMessagesSent sets the "messages_sent" field.
func (truo *TransactionReceiptUpdateOne) SetMessagesSent(t []*types.L1Message) *TransactionReceiptUpdateOne {
	truo.mutation.SetMessagesSent(t)
	return truo
}

// SetL1OriginMessage sets the "l1_origin_message" field.
func (truo *TransactionReceiptUpdateOne) SetL1OriginMessage(t *types.L2Message) *TransactionReceiptUpdateOne {
	truo.mutation.SetL1OriginMessage(t)
	return truo
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (truo *TransactionReceiptUpdateOne) SetTransactionID(id string) *TransactionReceiptUpdateOne {
	truo.mutation.SetTransactionID(id)
	return truo
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (truo *TransactionReceiptUpdateOne) SetNillableTransactionID(id *string) *TransactionReceiptUpdateOne {
	if id != nil {
		truo = truo.SetTransactionID(*id)
	}
	return truo
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (truo *TransactionReceiptUpdateOne) SetTransaction(t *Transaction) *TransactionReceiptUpdateOne {
	return truo.SetTransactionID(t.ID)
}

// Mutation returns the TransactionReceiptMutation object of the builder.
func (truo *TransactionReceiptUpdateOne) Mutation() *TransactionReceiptMutation {
	return truo.mutation
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (truo *TransactionReceiptUpdateOne) ClearTransaction() *TransactionReceiptUpdateOne {
	truo.mutation.ClearTransaction()
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TransactionReceiptUpdateOne) Select(field string, fields ...string) *TransactionReceiptUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TransactionReceipt entity.
func (truo *TransactionReceiptUpdateOne) Save(ctx context.Context) (*TransactionReceipt, error) {
	var (
		err  error
		node *TransactionReceipt
	)
	if len(truo.hooks) == 0 {
		if err = truo.check(); err != nil {
			return nil, err
		}
		node, err = truo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionReceiptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = truo.check(); err != nil {
				return nil, err
			}
			truo.mutation = mutation
			node, err = truo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(truo.hooks) - 1; i >= 0; i-- {
			if truo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = truo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, truo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TransactionReceiptUpdateOne) SaveX(ctx context.Context) *TransactionReceipt {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TransactionReceiptUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TransactionReceiptUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (truo *TransactionReceiptUpdateOne) check() error {
	if v, ok := truo.mutation.Status(); ok {
		if err := transactionreceipt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TransactionReceipt.status": %w`, err)}
		}
	}
	return nil
}

func (truo *TransactionReceiptUpdateOne) sqlSave(ctx context.Context) (_node *TransactionReceipt, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionreceipt.Table,
			Columns: transactionreceipt.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactionreceipt.FieldID,
			},
		},
	}
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TransactionReceipt.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transactionreceipt.FieldID)
		for _, f := range fields {
			if !transactionreceipt.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transactionreceipt.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.TransactionHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldTransactionHash,
		})
	}
	if value, ok := truo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transactionreceipt.FieldStatus,
		})
	}
	if value, ok := truo.mutation.StatusData(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldStatusData,
		})
	}
	if value, ok := truo.mutation.MessagesSent(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldMessagesSent,
		})
	}
	if value, ok := truo.mutation.L1OriginMessage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldL1OriginMessage,
		})
	}
	if truo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transactionreceipt.TransactionTable,
			Columns: []string{transactionreceipt.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   transactionreceipt.TransactionTable,
			Columns: []string{transactionreceipt.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TransactionReceipt{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionreceipt.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
