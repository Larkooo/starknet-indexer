// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tarrencev/starknet-indexer/ent/block"
	"github.com/tarrencev/starknet-indexer/ent/event"
	"github.com/tarrencev/starknet-indexer/ent/transaction"
	"github.com/tarrencev/starknet-indexer/ent/transactionreceipt"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
}

// SetContractAddress sets the "contract_address" field.
func (tc *TransactionCreate) SetContractAddress(s string) *TransactionCreate {
	tc.mutation.SetContractAddress(s)
	return tc
}

// SetEntryPointSelector sets the "entry_point_selector" field.
func (tc *TransactionCreate) SetEntryPointSelector(s string) *TransactionCreate {
	tc.mutation.SetEntryPointSelector(s)
	return tc
}

// SetNillableEntryPointSelector sets the "entry_point_selector" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableEntryPointSelector(s *string) *TransactionCreate {
	if s != nil {
		tc.SetEntryPointSelector(*s)
	}
	return tc
}

// SetTransactionHash sets the "transaction_hash" field.
func (tc *TransactionCreate) SetTransactionHash(s string) *TransactionCreate {
	tc.mutation.SetTransactionHash(s)
	return tc
}

// SetCalldata sets the "calldata" field.
func (tc *TransactionCreate) SetCalldata(s []string) *TransactionCreate {
	tc.mutation.SetCalldata(s)
	return tc
}

// SetSignature sets the "signature" field.
func (tc *TransactionCreate) SetSignature(s []string) *TransactionCreate {
	tc.mutation.SetSignature(s)
	return tc
}

// SetNonce sets the "nonce" field.
func (tc *TransactionCreate) SetNonce(s string) *TransactionCreate {
	tc.mutation.SetNonce(s)
	return tc
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableNonce(s *string) *TransactionCreate {
	if s != nil {
		tc.SetNonce(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(s string) *TransactionCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetBlockID sets the "block" edge to the Block entity by ID.
func (tc *TransactionCreate) SetBlockID(id string) *TransactionCreate {
	tc.mutation.SetBlockID(id)
	return tc
}

// SetNillableBlockID sets the "block" edge to the Block entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableBlockID(id *string) *TransactionCreate {
	if id != nil {
		tc = tc.SetBlockID(*id)
	}
	return tc
}

// SetBlock sets the "block" edge to the Block entity.
func (tc *TransactionCreate) SetBlock(b *Block) *TransactionCreate {
	return tc.SetBlockID(b.ID)
}

// SetReceiptID sets the "receipt" edge to the TransactionReceipt entity by ID.
func (tc *TransactionCreate) SetReceiptID(id string) *TransactionCreate {
	tc.mutation.SetReceiptID(id)
	return tc
}

// SetNillableReceiptID sets the "receipt" edge to the TransactionReceipt entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableReceiptID(id *string) *TransactionCreate {
	if id != nil {
		tc = tc.SetReceiptID(*id)
	}
	return tc
}

// SetReceipt sets the "receipt" edge to the TransactionReceipt entity.
func (tc *TransactionCreate) SetReceipt(t *TransactionReceipt) *TransactionCreate {
	return tc.SetReceiptID(t.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (tc *TransactionCreate) AddEventIDs(ids ...string) *TransactionCreate {
	tc.mutation.AddEventIDs(ids...)
	return tc
}

// AddEvents adds the "events" edges to the Event entity.
func (tc *TransactionCreate) AddEvents(e ...*Event) *TransactionCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tc.AddEventIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.ContractAddress(); !ok {
		return &ValidationError{Name: "contract_address", err: errors.New(`ent: missing required field "Transaction.contract_address"`)}
	}
	if _, ok := tc.mutation.TransactionHash(); !ok {
		return &ValidationError{Name: "transaction_hash", err: errors.New(`ent: missing required field "Transaction.transaction_hash"`)}
	}
	if _, ok := tc.mutation.Calldata(); !ok {
		return &ValidationError{Name: "calldata", err: errors.New(`ent: missing required field "Transaction.calldata"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Transaction.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transaction.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.ContractAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldContractAddress,
		})
		_node.ContractAddress = value
	}
	if value, ok := tc.mutation.EntryPointSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldEntryPointSelector,
		})
		_node.EntryPointSelector = value
	}
	if value, ok := tc.mutation.TransactionHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionHash,
		})
		_node.TransactionHash = value
	}
	if value, ok := tc.mutation.Calldata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldCalldata,
		})
		_node.Calldata = value
	}
	if value, ok := tc.mutation.Signature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldSignature,
		})
		_node.Signature = value
	}
	if value, ok := tc.mutation.Nonce(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldNonce,
		})
		_node.Nonce = value
	}
	if nodes := tc.mutation.BlockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.BlockTable,
			Columns: []string{transaction.BlockColumn},
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
		_node.block_transactions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transaction.ReceiptTable,
			Columns: []string{transaction.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transactionreceipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.EventsTable,
			Columns: []string{transaction.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
