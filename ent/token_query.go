// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/contract"
	"github.com/cartridge-gg/starknet-indexer/ent/predicate"
	"github.com/cartridge-gg/starknet-indexer/ent/token"
)

// TokenQuery is the builder for querying Token entities.
type TokenQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Token
	// eager-loading edges.
	withOwner    *ContractQuery
	withContract *ContractQuery
	withFKs      bool
	modifiers    []func(*sql.Selector)
	loadTotal    []func(context.Context, []*Token) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TokenQuery builder.
func (tq *TokenQuery) Where(ps ...predicate.Token) *TokenQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TokenQuery) Limit(limit int) *TokenQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TokenQuery) Offset(offset int) *TokenQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TokenQuery) Unique(unique bool) *TokenQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TokenQuery) Order(o ...OrderFunc) *TokenQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryOwner chains the current query on the "owner" edge.
func (tq *TokenQuery) QueryOwner() *ContractQuery {
	query := &ContractQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, selector),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, token.OwnerTable, token.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryContract chains the current query on the "contract" edge.
func (tq *TokenQuery) QueryContract() *ContractQuery {
	query := &ContractQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, selector),
			sqlgraph.To(contract.Table, contract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, token.ContractTable, token.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Token entity from the query.
// Returns a *NotFoundError when no Token was found.
func (tq *TokenQuery) First(ctx context.Context) (*Token, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{token.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TokenQuery) FirstX(ctx context.Context) *Token {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Token ID from the query.
// Returns a *NotFoundError when no Token ID was found.
func (tq *TokenQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{token.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TokenQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Token entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Token entity is found.
// Returns a *NotFoundError when no Token entities are found.
func (tq *TokenQuery) Only(ctx context.Context) (*Token, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{token.Label}
	default:
		return nil, &NotSingularError{token.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TokenQuery) OnlyX(ctx context.Context) *Token {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Token ID in the query.
// Returns a *NotSingularError when more than one Token ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TokenQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{token.Label}
	default:
		err = &NotSingularError{token.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TokenQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tokens.
func (tq *TokenQuery) All(ctx context.Context) ([]*Token, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TokenQuery) AllX(ctx context.Context) []*Token {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Token IDs.
func (tq *TokenQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := tq.Select(token.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TokenQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TokenQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TokenQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TokenQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TokenQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TokenQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TokenQuery) Clone() *TokenQuery {
	if tq == nil {
		return nil
	}
	return &TokenQuery{
		config:       tq.config,
		limit:        tq.limit,
		offset:       tq.offset,
		order:        append([]OrderFunc{}, tq.order...),
		predicates:   append([]predicate.Token{}, tq.predicates...),
		withOwner:    tq.withOwner.Clone(),
		withContract: tq.withContract.Clone(),
		// clone intermediate query.
		sql:    tq.sql.Clone(),
		path:   tq.path,
		unique: tq.unique,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TokenQuery) WithOwner(opts ...func(*ContractQuery)) *TokenQuery {
	query := &ContractQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withOwner = query
	return tq
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TokenQuery) WithContract(opts ...func(*ContractQuery)) *TokenQuery {
	query := &ContractQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withContract = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TokenId uint64 `json:"tokenId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Token.Query().
//		GroupBy(token.FieldTokenId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TokenQuery) GroupBy(field string, fields ...string) *TokenGroupBy {
	grbuild := &TokenGroupBy{config: tq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	grbuild.label = token.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TokenId uint64 `json:"tokenId,omitempty"`
//	}
//
//	client.Token.Query().
//		Select(token.FieldTokenId).
//		Scan(ctx, &v)
//
func (tq *TokenQuery) Select(fields ...string) *TokenSelect {
	tq.fields = append(tq.fields, fields...)
	selbuild := &TokenSelect{TokenQuery: tq}
	selbuild.label = token.Label
	selbuild.flds, selbuild.scan = &tq.fields, selbuild.Scan
	return selbuild
}

func (tq *TokenQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !token.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TokenQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Token, error) {
	var (
		nodes       = []*Token{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withOwner != nil,
			tq.withContract != nil,
		}
	)
	if tq.withOwner != nil || tq.withContract != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, token.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Token).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Token{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withOwner; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Token)
		for i := range nodes {
			if nodes[i].token_owner == nil {
				continue
			}
			fk := *nodes[i].token_owner
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(contract.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "token_owner" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Owner = n
			}
		}
	}

	if query := tq.withContract; query != nil {
		ids := make([]string, 0, len(nodes))
		nodeids := make(map[string][]*Token)
		for i := range nodes {
			if nodes[i].token_contract == nil {
				continue
			}
			fk := *nodes[i].token_contract
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(contract.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "token_contract" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Contract = n
			}
		}
	}

	for i := range tq.loadTotal {
		if err := tq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TokenQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	if len(tq.modifiers) > 0 {
		_spec.Modifiers = tq.modifiers
	}
	_spec.Node.Columns = tq.fields
	if len(tq.fields) > 0 {
		_spec.Unique = tq.unique != nil && *tq.unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TokenQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TokenQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: token.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for i := range fields {
			if fields[i] != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TokenQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(token.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = token.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.unique != nil && *tq.unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TokenGroupBy is the group-by builder for Token entities.
type TokenGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TokenGroupBy) Aggregate(fns ...AggregateFunc) *TokenGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TokenGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

func (tgb *TokenGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !token.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TokenGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TokenSelect is the builder for selecting fields of Token entities.
type TokenSelect struct {
	*TokenQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TokenSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TokenQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

func (ts *TokenSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
