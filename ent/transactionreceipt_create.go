// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dontpanicdao/caigo/types"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/transaction"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

// TransactionReceiptCreate is the builder for creating a TransactionReceipt entity.
type TransactionReceiptCreate struct {
	config
	mutation *TransactionReceiptMutation
	hooks    []Hook
}

// SetTransactionHash sets the "transaction_hash" field.
func (trc *TransactionReceiptCreate) SetTransactionHash(s string) *TransactionReceiptCreate {
	trc.mutation.SetTransactionHash(s)
	return trc
}

// SetStatus sets the "status" field.
func (trc *TransactionReceiptCreate) SetStatus(t transactionreceipt.Status) *TransactionReceiptCreate {
	trc.mutation.SetStatus(t)
	return trc
}

// SetStatusData sets the "status_data" field.
func (trc *TransactionReceiptCreate) SetStatusData(s string) *TransactionReceiptCreate {
	trc.mutation.SetStatusData(s)
	return trc
}

// SetMessagesSent sets the "messages_sent" field.
func (trc *TransactionReceiptCreate) SetMessagesSent(t []*types.L1Message) *TransactionReceiptCreate {
	trc.mutation.SetMessagesSent(t)
	return trc
}

// SetL1OriginMessage sets the "l1_origin_message" field.
func (trc *TransactionReceiptCreate) SetL1OriginMessage(t *types.L2Message) *TransactionReceiptCreate {
	trc.mutation.SetL1OriginMessage(t)
	return trc
}

// SetID sets the "id" field.
func (trc *TransactionReceiptCreate) SetID(s string) *TransactionReceiptCreate {
	trc.mutation.SetID(s)
	return trc
}

// SetBlockID sets the "block" edge to the Block entity by ID.
func (trc *TransactionReceiptCreate) SetBlockID(id string) *TransactionReceiptCreate {
	trc.mutation.SetBlockID(id)
	return trc
}

// SetNillableBlockID sets the "block" edge to the Block entity by ID if the given value is not nil.
func (trc *TransactionReceiptCreate) SetNillableBlockID(id *string) *TransactionReceiptCreate {
	if id != nil {
		trc = trc.SetBlockID(*id)
	}
	return trc
}

// SetBlock sets the "block" edge to the Block entity.
func (trc *TransactionReceiptCreate) SetBlock(b *Block) *TransactionReceiptCreate {
	return trc.SetBlockID(b.ID)
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (trc *TransactionReceiptCreate) SetTransactionID(id string) *TransactionReceiptCreate {
	trc.mutation.SetTransactionID(id)
	return trc
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (trc *TransactionReceiptCreate) SetTransaction(t *Transaction) *TransactionReceiptCreate {
	return trc.SetTransactionID(t.ID)
}

// Mutation returns the TransactionReceiptMutation object of the builder.
func (trc *TransactionReceiptCreate) Mutation() *TransactionReceiptMutation {
	return trc.mutation
}

// Save creates the TransactionReceipt in the database.
func (trc *TransactionReceiptCreate) Save(ctx context.Context) (*TransactionReceipt, error) {
	var (
		err  error
		node *TransactionReceipt
	)
	if len(trc.hooks) == 0 {
		if err = trc.check(); err != nil {
			return nil, err
		}
		node, err = trc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionReceiptMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = trc.check(); err != nil {
				return nil, err
			}
			trc.mutation = mutation
			if node, err = trc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(trc.hooks) - 1; i >= 0; i-- {
			if trc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = trc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, trc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (trc *TransactionReceiptCreate) SaveX(ctx context.Context) *TransactionReceipt {
	v, err := trc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trc *TransactionReceiptCreate) Exec(ctx context.Context) error {
	_, err := trc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trc *TransactionReceiptCreate) ExecX(ctx context.Context) {
	if err := trc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (trc *TransactionReceiptCreate) check() error {
	if _, ok := trc.mutation.TransactionHash(); !ok {
		return &ValidationError{Name: "transaction_hash", err: errors.New(`ent: missing required field "TransactionReceipt.transaction_hash"`)}
	}
	if _, ok := trc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "TransactionReceipt.status"`)}
	}
	if v, ok := trc.mutation.Status(); ok {
		if err := transactionreceipt.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "TransactionReceipt.status": %w`, err)}
		}
	}
	if _, ok := trc.mutation.StatusData(); !ok {
		return &ValidationError{Name: "status_data", err: errors.New(`ent: missing required field "TransactionReceipt.status_data"`)}
	}
	if _, ok := trc.mutation.MessagesSent(); !ok {
		return &ValidationError{Name: "messages_sent", err: errors.New(`ent: missing required field "TransactionReceipt.messages_sent"`)}
	}
	if _, ok := trc.mutation.L1OriginMessage(); !ok {
		return &ValidationError{Name: "l1_origin_message", err: errors.New(`ent: missing required field "TransactionReceipt.l1_origin_message"`)}
	}
	if _, ok := trc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction", err: errors.New(`ent: missing required edge "TransactionReceipt.transaction"`)}
	}
	return nil
}

func (trc *TransactionReceiptCreate) sqlSave(ctx context.Context) (*TransactionReceipt, error) {
	_node, _spec := trc.createSpec()
	if err := sqlgraph.CreateNode(ctx, trc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected TransactionReceipt.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (trc *TransactionReceiptCreate) createSpec() (*TransactionReceipt, *sqlgraph.CreateSpec) {
	var (
		_node = &TransactionReceipt{config: trc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transactionreceipt.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transactionreceipt.FieldID,
			},
		}
	)
	if id, ok := trc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := trc.mutation.TransactionHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldTransactionHash,
		})
		_node.TransactionHash = value
	}
	if value, ok := trc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: transactionreceipt.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := trc.mutation.StatusData(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionreceipt.FieldStatusData,
		})
		_node.StatusData = value
	}
	if value, ok := trc.mutation.MessagesSent(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldMessagesSent,
		})
		_node.MessagesSent = value
	}
	if value, ok := trc.mutation.L1OriginMessage(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transactionreceipt.FieldL1OriginMessage,
		})
		_node.L1OriginMessage = value
	}
	if nodes := trc.mutation.BlockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionreceipt.BlockTable,
			Columns: []string{transactionreceipt.BlockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: block.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.block_transaction_receipts = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := trc.mutation.TransactionIDs(); len(nodes) > 0 {
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
		_node.transaction_receipt = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionReceiptCreateBulk is the builder for creating many TransactionReceipt entities in bulk.
type TransactionReceiptCreateBulk struct {
	config
	builders []*TransactionReceiptCreate
}

// Save creates the TransactionReceipt entities in the database.
func (trcb *TransactionReceiptCreateBulk) Save(ctx context.Context) ([]*TransactionReceipt, error) {
	specs := make([]*sqlgraph.CreateSpec, len(trcb.builders))
	nodes := make([]*TransactionReceipt, len(trcb.builders))
	mutators := make([]Mutator, len(trcb.builders))
	for i := range trcb.builders {
		func(i int, root context.Context) {
			builder := trcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionReceiptMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, trcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, trcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, trcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (trcb *TransactionReceiptCreateBulk) SaveX(ctx context.Context) []*TransactionReceipt {
	v, err := trcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (trcb *TransactionReceiptCreateBulk) Exec(ctx context.Context) error {
	_, err := trcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (trcb *TransactionReceiptCreateBulk) ExecX(ctx context.Context) {
	if err := trcb.Exec(ctx); err != nil {
		panic(err)
	}
}
