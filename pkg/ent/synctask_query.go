// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"fsm/pkg/ent/predicate"
	"fsm/pkg/ent/synctask"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SyncTaskQuery is the builder for querying SyncTask entities.
type SyncTaskQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SyncTask
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SyncTaskQuery builder.
func (stq *SyncTaskQuery) Where(ps ...predicate.SyncTask) *SyncTaskQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit adds a limit step to the query.
func (stq *SyncTaskQuery) Limit(limit int) *SyncTaskQuery {
	stq.limit = &limit
	return stq
}

// Offset adds an offset step to the query.
func (stq *SyncTaskQuery) Offset(offset int) *SyncTaskQuery {
	stq.offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *SyncTaskQuery) Unique(unique bool) *SyncTaskQuery {
	stq.unique = &unique
	return stq
}

// Order adds an order step to the query.
func (stq *SyncTaskQuery) Order(o ...OrderFunc) *SyncTaskQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// First returns the first SyncTask entity from the query.
// Returns a *NotFoundError when no SyncTask was found.
func (stq *SyncTaskQuery) First(ctx context.Context) (*SyncTask, error) {
	nodes, err := stq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{synctask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *SyncTaskQuery) FirstX(ctx context.Context) *SyncTask {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SyncTask ID from the query.
// Returns a *NotFoundError when no SyncTask ID was found.
func (stq *SyncTaskQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = stq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{synctask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *SyncTaskQuery) FirstIDX(ctx context.Context) string {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SyncTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SyncTask entity is found.
// Returns a *NotFoundError when no SyncTask entities are found.
func (stq *SyncTaskQuery) Only(ctx context.Context) (*SyncTask, error) {
	nodes, err := stq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{synctask.Label}
	default:
		return nil, &NotSingularError{synctask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *SyncTaskQuery) OnlyX(ctx context.Context) *SyncTask {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SyncTask ID in the query.
// Returns a *NotSingularError when more than one SyncTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *SyncTaskQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = stq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{synctask.Label}
	default:
		err = &NotSingularError{synctask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *SyncTaskQuery) OnlyIDX(ctx context.Context) string {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SyncTasks.
func (stq *SyncTaskQuery) All(ctx context.Context) ([]*SyncTask, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return stq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (stq *SyncTaskQuery) AllX(ctx context.Context) []*SyncTask {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SyncTask IDs.
func (stq *SyncTaskQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := stq.Select(synctask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *SyncTaskQuery) IDsX(ctx context.Context) []string {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *SyncTaskQuery) Count(ctx context.Context) (int, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return stq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (stq *SyncTaskQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *SyncTaskQuery) Exist(ctx context.Context) (bool, error) {
	if err := stq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return stq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *SyncTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SyncTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *SyncTaskQuery) Clone() *SyncTaskQuery {
	if stq == nil {
		return nil
	}
	return &SyncTaskQuery{
		config:     stq.config,
		limit:      stq.limit,
		offset:     stq.offset,
		order:      append([]OrderFunc{}, stq.order...),
		predicates: append([]predicate.SyncTask{}, stq.predicates...),
		// clone intermediate query.
		sql:    stq.sql.Clone(),
		path:   stq.path,
		unique: stq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SyncTask.Query().
//		GroupBy(synctask.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (stq *SyncTaskQuery) GroupBy(field string, fields ...string) *SyncTaskGroupBy {
	grbuild := &SyncTaskGroupBy{config: stq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return stq.sqlQuery(ctx), nil
	}
	grbuild.label = synctask.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.SyncTask.Query().
//		Select(synctask.FieldUserID).
//		Scan(ctx, &v)
func (stq *SyncTaskQuery) Select(fields ...string) *SyncTaskSelect {
	stq.fields = append(stq.fields, fields...)
	selbuild := &SyncTaskSelect{SyncTaskQuery: stq}
	selbuild.label = synctask.Label
	selbuild.flds, selbuild.scan = &stq.fields, selbuild.Scan
	return selbuild
}

func (stq *SyncTaskQuery) prepareQuery(ctx context.Context) error {
	for _, f := range stq.fields {
		if !synctask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *SyncTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SyncTask, error) {
	var (
		nodes = []*SyncTask{}
		_spec = stq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SyncTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SyncTask{config: stq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (stq *SyncTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.fields
	if len(stq.fields) > 0 {
		_spec.Unique = stq.unique != nil && *stq.unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *SyncTaskQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (stq *SyncTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   synctask.Table,
			Columns: synctask.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: synctask.FieldID,
			},
		},
		From:   stq.sql,
		Unique: true,
	}
	if unique := stq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := stq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, synctask.FieldID)
		for i := range fields {
			if fields[i] != synctask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *SyncTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(synctask.Table)
	columns := stq.fields
	if len(columns) == 0 {
		columns = synctask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.unique != nil && *stq.unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SyncTaskGroupBy is the group-by builder for SyncTask entities.
type SyncTaskGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *SyncTaskGroupBy) Aggregate(fns ...AggregateFunc) *SyncTaskGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the group-by query and scans the result into the given value.
func (stgb *SyncTaskGroupBy) Scan(ctx context.Context, v any) error {
	query, err := stgb.path(ctx)
	if err != nil {
		return err
	}
	stgb.sql = query
	return stgb.sqlScan(ctx, v)
}

func (stgb *SyncTaskGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range stgb.fields {
		if !synctask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := stgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (stgb *SyncTaskGroupBy) sqlQuery() *sql.Selector {
	selector := stgb.sql.Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(stgb.fields)+len(stgb.fns))
		for _, f := range stgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(stgb.fields...)...)
}

// SyncTaskSelect is the builder for selecting fields of SyncTask entities.
type SyncTaskSelect struct {
	*SyncTaskQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sts *SyncTaskSelect) Scan(ctx context.Context, v any) error {
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	sts.sql = sts.SyncTaskQuery.sqlQuery(ctx)
	return sts.sqlScan(ctx, v)
}

func (sts *SyncTaskSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sts.sql.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
