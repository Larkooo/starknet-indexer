// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/predicate"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
)

// BlockUpdate is the builder for updating Block entities.
type BlockUpdate struct {
	config
	hooks    []Hook
	mutation *BlockMutation
}

// Where appends a list predicates to the BlockUpdate builder.
func (bu *BlockUpdate) Where(ps ...predicate.Block) *BlockUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetBlockHash sets the "block_hash" field.
func (bu *BlockUpdate) SetBlockHash(s string) *BlockUpdate {
	bu.mutation.SetBlockHash(s)
	return bu
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (bu *BlockUpdate) SetParentBlockHash(s string) *BlockUpdate {
	bu.mutation.SetParentBlockHash(s)
	return bu
}

// SetBlockNumber sets the "block_number" field.
func (bu *BlockUpdate) SetBlockNumber(u uint64) *BlockUpdate {
	bu.mutation.ResetBlockNumber()
	bu.mutation.SetBlockNumber(u)
	return bu
}

// AddBlockNumber adds u to the "block_number" field.
func (bu *BlockUpdate) AddBlockNumber(u int64) *BlockUpdate {
	bu.mutation.AddBlockNumber(u)
	return bu
}

// SetStateRoot sets the "state_root" field.
func (bu *BlockUpdate) SetStateRoot(s string) *BlockUpdate {
	bu.mutation.SetStateRoot(s)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BlockUpdate) SetStatus(b block.Status) *BlockUpdate {
	bu.mutation.SetStatus(b)
	return bu
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (bu *BlockUpdate) AddTransactionIDs(ids ...string) *BlockUpdate {
	bu.mutation.AddTransactionIDs(ids...)
	return bu
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (bu *BlockUpdate) AddTransactions(t ...*Transaction) *BlockUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTransactionIDs(ids...)
}

// AddTransactionReceiptIDs adds the "transaction_receipts" edge to the TransactionReceipt entity by IDs.
func (bu *BlockUpdate) AddTransactionReceiptIDs(ids ...string) *BlockUpdate {
	bu.mutation.AddTransactionReceiptIDs(ids...)
	return bu
}

// AddTransactionReceipts adds the "transaction_receipts" edges to the TransactionReceipt entity.
func (bu *BlockUpdate) AddTransactionReceipts(t ...*TransactionReceipt) *BlockUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTransactionReceiptIDs(ids...)
}

// Mutation returns the BlockMutation object of the builder.
func (bu *BlockUpdate) Mutation() *BlockMutation {
	return bu.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (bu *BlockUpdate) ClearTransactions() *BlockUpdate {
	bu.mutation.ClearTransactions()
	return bu
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (bu *BlockUpdate) RemoveTransactionIDs(ids ...string) *BlockUpdate {
	bu.mutation.RemoveTransactionIDs(ids...)
	return bu
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (bu *BlockUpdate) RemoveTransactions(t ...*Transaction) *BlockUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTransactionIDs(ids...)
}

// ClearTransactionReceipts clears all "transaction_receipts" edges to the TransactionReceipt entity.
func (bu *BlockUpdate) ClearTransactionReceipts() *BlockUpdate {
	bu.mutation.ClearTransactionReceipts()
	return bu
}

// RemoveTransactionReceiptIDs removes the "transaction_receipts" edge to TransactionReceipt entities by IDs.
func (bu *BlockUpdate) RemoveTransactionReceiptIDs(ids ...string) *BlockUpdate {
	bu.mutation.RemoveTransactionReceiptIDs(ids...)
	return bu
}

// RemoveTransactionReceipts removes "transaction_receipts" edges to TransactionReceipt entities.
func (bu *BlockUpdate) RemoveTransactionReceipts(t ...*TransactionReceipt) *BlockUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTransactionReceiptIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		if err = bu.check(); err != nil {
			return 0, err
		}
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bu.check(); err != nil {
				return 0, err
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			if bu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlockUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlockUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlockUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BlockUpdate) check() error {
	if v, ok := bu.mutation.Status(); ok {
		if err := block.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Block.status": %w`, err)}
		}
	}
	return nil
}

func (bu *BlockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: block.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldBlockHash,
		})
	}
	if value, ok := bu.mutation.ParentBlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldParentBlockHash,
		})
	}
	if value, ok := bu.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := bu.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := bu.mutation.StateRoot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldStateRoot,
		})
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: block.FieldStatus,
		})
	}
	if bu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
	if nodes := bu.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !bu.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
	if bu.mutation.TransactionReceiptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transactionreceipt.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTransactionReceiptsIDs(); len(nodes) > 0 && !bu.mutation.TransactionReceiptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TransactionReceiptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// BlockUpdateOne is the builder for updating a single Block entity.
type BlockUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlockMutation
}

// SetBlockHash sets the "block_hash" field.
func (buo *BlockUpdateOne) SetBlockHash(s string) *BlockUpdateOne {
	buo.mutation.SetBlockHash(s)
	return buo
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (buo *BlockUpdateOne) SetParentBlockHash(s string) *BlockUpdateOne {
	buo.mutation.SetParentBlockHash(s)
	return buo
}

// SetBlockNumber sets the "block_number" field.
func (buo *BlockUpdateOne) SetBlockNumber(u uint64) *BlockUpdateOne {
	buo.mutation.ResetBlockNumber()
	buo.mutation.SetBlockNumber(u)
	return buo
}

// AddBlockNumber adds u to the "block_number" field.
func (buo *BlockUpdateOne) AddBlockNumber(u int64) *BlockUpdateOne {
	buo.mutation.AddBlockNumber(u)
	return buo
}

// SetStateRoot sets the "state_root" field.
func (buo *BlockUpdateOne) SetStateRoot(s string) *BlockUpdateOne {
	buo.mutation.SetStateRoot(s)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BlockUpdateOne) SetStatus(b block.Status) *BlockUpdateOne {
	buo.mutation.SetStatus(b)
	return buo
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (buo *BlockUpdateOne) AddTransactionIDs(ids ...string) *BlockUpdateOne {
	buo.mutation.AddTransactionIDs(ids...)
	return buo
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (buo *BlockUpdateOne) AddTransactions(t ...*Transaction) *BlockUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTransactionIDs(ids...)
}

// AddTransactionReceiptIDs adds the "transaction_receipts" edge to the TransactionReceipt entity by IDs.
func (buo *BlockUpdateOne) AddTransactionReceiptIDs(ids ...string) *BlockUpdateOne {
	buo.mutation.AddTransactionReceiptIDs(ids...)
	return buo
}

// AddTransactionReceipts adds the "transaction_receipts" edges to the TransactionReceipt entity.
func (buo *BlockUpdateOne) AddTransactionReceipts(t ...*TransactionReceipt) *BlockUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTransactionReceiptIDs(ids...)
}

// Mutation returns the BlockMutation object of the builder.
func (buo *BlockUpdateOne) Mutation() *BlockMutation {
	return buo.mutation
}

// ClearTransactions clears all "transactions" edges to the Transaction entity.
func (buo *BlockUpdateOne) ClearTransactions() *BlockUpdateOne {
	buo.mutation.ClearTransactions()
	return buo
}

// RemoveTransactionIDs removes the "transactions" edge to Transaction entities by IDs.
func (buo *BlockUpdateOne) RemoveTransactionIDs(ids ...string) *BlockUpdateOne {
	buo.mutation.RemoveTransactionIDs(ids...)
	return buo
}

// RemoveTransactions removes "transactions" edges to Transaction entities.
func (buo *BlockUpdateOne) RemoveTransactions(t ...*Transaction) *BlockUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTransactionIDs(ids...)
}

// ClearTransactionReceipts clears all "transaction_receipts" edges to the TransactionReceipt entity.
func (buo *BlockUpdateOne) ClearTransactionReceipts() *BlockUpdateOne {
	buo.mutation.ClearTransactionReceipts()
	return buo
}

// RemoveTransactionReceiptIDs removes the "transaction_receipts" edge to TransactionReceipt entities by IDs.
func (buo *BlockUpdateOne) RemoveTransactionReceiptIDs(ids ...string) *BlockUpdateOne {
	buo.mutation.RemoveTransactionReceiptIDs(ids...)
	return buo
}

// RemoveTransactionReceipts removes "transaction_receipts" edges to TransactionReceipt entities.
func (buo *BlockUpdateOne) RemoveTransactionReceipts(t ...*TransactionReceipt) *BlockUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTransactionReceiptIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlockUpdateOne) Select(field string, fields ...string) *BlockUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Block entity.
func (buo *BlockUpdateOne) Save(ctx context.Context) (*Block, error) {
	var (
		err  error
		node *Block
	)
	if len(buo.hooks) == 0 {
		if err = buo.check(); err != nil {
			return nil, err
		}
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = buo.check(); err != nil {
				return nil, err
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			if buo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = buo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, buo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Block)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlockUpdateOne) SaveX(ctx context.Context) *Block {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlockUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlockUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BlockUpdateOne) check() error {
	if v, ok := buo.mutation.Status(); ok {
		if err := block.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Block.status": %w`, err)}
		}
	}
	return nil
}

func (buo *BlockUpdateOne) sqlSave(ctx context.Context) (_node *Block, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   block.Table,
			Columns: block.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: block.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Block.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, block.FieldID)
		for _, f := range fields {
			if !block.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != block.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.BlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldBlockHash,
		})
	}
	if value, ok := buo.mutation.ParentBlockHash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldParentBlockHash,
		})
	}
	if value, ok := buo.mutation.BlockNumber(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := buo.mutation.AddedBlockNumber(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
	}
	if value, ok := buo.mutation.StateRoot(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldStateRoot,
		})
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: block.FieldStatus,
		})
	}
	if buo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
	if nodes := buo.mutation.RemovedTransactionsIDs(); len(nodes) > 0 && !buo.mutation.TransactionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
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
	if buo.mutation.TransactionReceiptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transactionreceipt.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTransactionReceiptsIDs(); len(nodes) > 0 && !buo.mutation.TransactionReceiptsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TransactionReceiptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Block{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{block.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
