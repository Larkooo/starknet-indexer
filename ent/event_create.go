// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/event"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/dontpanicdao/caigo/types"
)

// EventCreate is the builder for creating a Event entity.
type EventCreate struct {
	config
	mutation *EventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetFrom sets the "from" field.
func (ec *EventCreate) SetFrom(s string) *EventCreate {
	ec.mutation.SetFrom(s)
	return ec
}

// SetKeys sets the "keys" field.
func (ec *EventCreate) SetKeys(t []*types.Felt) *EventCreate {
	ec.mutation.SetKeys(t)
	return ec
}

// SetData sets the "data" field.
func (ec *EventCreate) SetData(t []*types.Felt) *EventCreate {
	ec.mutation.SetData(t)
	return ec
}

// SetID sets the "id" field.
func (ec *EventCreate) SetID(s string) *EventCreate {
	ec.mutation.SetID(s)
	return ec
}

// SetTransactionID sets the "transaction" edge to the Transaction entity by ID.
func (ec *EventCreate) SetTransactionID(id string) *EventCreate {
	ec.mutation.SetTransactionID(id)
	return ec
}

// SetNillableTransactionID sets the "transaction" edge to the Transaction entity by ID if the given value is not nil.
func (ec *EventCreate) SetNillableTransactionID(id *string) *EventCreate {
	if id != nil {
		ec = ec.SetTransactionID(*id)
	}
	return ec
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (ec *EventCreate) SetTransaction(t *Transaction) *EventCreate {
	return ec.SetTransactionID(t.ID)
}

// Mutation returns the EventMutation object of the builder.
func (ec *EventCreate) Mutation() *EventMutation {
	return ec.mutation
}

// Save creates the Event in the database.
func (ec *EventCreate) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Event)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EventMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EventCreate) SaveX(ctx context.Context) *Event {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EventCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EventCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EventCreate) check() error {
	if _, ok := ec.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "Event.from"`)}
	}
	if _, ok := ec.mutation.Keys(); !ok {
		return &ValidationError{Name: "keys", err: errors.New(`ent: missing required field "Event.keys"`)}
	}
	if _, ok := ec.mutation.Data(); !ok {
		return &ValidationError{Name: "data", err: errors.New(`ent: missing required field "Event.data"`)}
	}
	return nil
}

func (ec *EventCreate) sqlSave(ctx context.Context) (*Event, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Event.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (ec *EventCreate) createSpec() (*Event, *sqlgraph.CreateSpec) {
	var (
		_node = &Event{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: event.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: event.FieldID,
			},
		}
	)
	_spec.OnConflict = ec.conflict
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.From(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldFrom,
		})
		_node.From = value
	}
	if value, ok := ec.mutation.Keys(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldKeys,
		})
		_node.Keys = value
	}
	if value, ok := ec.mutation.Data(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: event.FieldData,
		})
		_node.Data = value
	}
	if nodes := ec.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.TransactionTable,
			Columns: []string{event.TransactionColumn},
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
		_node.transaction_events = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.Create().
//		SetFrom(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetFrom(v+v).
//		}).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflict(opts ...sql.ConflictOption) *EventUpsertOne {
	ec.conflict = opts
	return &EventUpsertOne{
		create: ec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ec *EventCreate) OnConflictColumns(columns ...string) *EventUpsertOne {
	ec.conflict = append(ec.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertOne{
		create: ec,
	}
}

type (
	// EventUpsertOne is the builder for "upsert"-ing
	//  one Event node.
	EventUpsertOne struct {
		create *EventCreate
	}

	// EventUpsert is the "OnConflict" setter.
	EventUpsert struct {
		*sql.UpdateSet
	}
)

// SetFrom sets the "from" field.
func (u *EventUpsert) SetFrom(v string) *EventUpsert {
	u.Set(event.FieldFrom, v)
	return u
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EventUpsert) UpdateFrom() *EventUpsert {
	u.SetExcluded(event.FieldFrom)
	return u
}

// SetKeys sets the "keys" field.
func (u *EventUpsert) SetKeys(v []*types.Felt) *EventUpsert {
	u.Set(event.FieldKeys, v)
	return u
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *EventUpsert) UpdateKeys() *EventUpsert {
	u.SetExcluded(event.FieldKeys)
	return u
}

// SetData sets the "data" field.
func (u *EventUpsert) SetData(v []*types.Felt) *EventUpsert {
	u.Set(event.FieldData, v)
	return u
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *EventUpsert) UpdateData() *EventUpsert {
	u.SetExcluded(event.FieldData)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(event.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *EventUpsertOne) UpdateNewValues() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(event.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Event.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *EventUpsertOne) Ignore() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertOne) DoNothing() *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreate.OnConflict
// documentation for more info.
func (u *EventUpsertOne) Update(set func(*EventUpsert)) *EventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetFrom sets the "from" field.
func (u *EventUpsertOne) SetFrom(v string) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateFrom() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateFrom()
	})
}

// SetKeys sets the "keys" field.
func (u *EventUpsertOne) SetKeys(v []*types.Felt) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetKeys(v)
	})
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateKeys() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateKeys()
	})
}

// SetData sets the "data" field.
func (u *EventUpsertOne) SetData(v []*types.Felt) *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *EventUpsertOne) UpdateData() *EventUpsertOne {
	return u.Update(func(s *EventUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *EventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *EventUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: EventUpsertOne.ID is not supported by MySQL driver. Use EventUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *EventUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// EventCreateBulk is the builder for creating many Event entities in bulk.
type EventCreateBulk struct {
	config
	builders []*EventCreate
	conflict []sql.ConflictOption
}

// Save creates the Event entities in the database.
func (ecb *EventCreateBulk) Save(ctx context.Context) ([]*Event, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Event, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EventCreateBulk) SaveX(ctx context.Context) []*Event {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EventCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EventCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Event.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.EventUpsert) {
//			SetFrom(v+v).
//		}).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflict(opts ...sql.ConflictOption) *EventUpsertBulk {
	ecb.conflict = opts
	return &EventUpsertBulk{
		create: ecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ecb *EventCreateBulk) OnConflictColumns(columns ...string) *EventUpsertBulk {
	ecb.conflict = append(ecb.conflict, sql.ConflictColumns(columns...))
	return &EventUpsertBulk{
		create: ecb,
	}
}

// EventUpsertBulk is the builder for "upsert"-ing
// a bulk of Event nodes.
type EventUpsertBulk struct {
	create *EventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(event.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *EventUpsertBulk) UpdateNewValues() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(event.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Event.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *EventUpsertBulk) Ignore() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *EventUpsertBulk) DoNothing() *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the EventCreateBulk.OnConflict
// documentation for more info.
func (u *EventUpsertBulk) Update(set func(*EventUpsert)) *EventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&EventUpsert{UpdateSet: update})
	}))
	return u
}

// SetFrom sets the "from" field.
func (u *EventUpsertBulk) SetFrom(v string) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetFrom(v)
	})
}

// UpdateFrom sets the "from" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateFrom() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateFrom()
	})
}

// SetKeys sets the "keys" field.
func (u *EventUpsertBulk) SetKeys(v []*types.Felt) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetKeys(v)
	})
}

// UpdateKeys sets the "keys" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateKeys() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateKeys()
	})
}

// SetData sets the "data" field.
func (u *EventUpsertBulk) SetData(v []*types.Felt) *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.SetData(v)
	})
}

// UpdateData sets the "data" field to the value that was provided on create.
func (u *EventUpsertBulk) UpdateData() *EventUpsertBulk {
	return u.Update(func(s *EventUpsert) {
		s.UpdateData()
	})
}

// Exec executes the query.
func (u *EventUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the EventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for EventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *EventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
