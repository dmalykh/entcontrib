// Code generated by ent, DO NOT EDIT.

package oastypes

import (
	"context"
	"fmt"
	"math"

	"github.com/dmalykh/entcontrib/entoas/internal/oastypes/oastypes"
	"github.com/dmalykh/entcontrib/entoas/internal/oastypes/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OASTypesQuery is the builder for querying OASTypes entities.
type OASTypesQuery struct {
	config
	ctx        *QueryContext
	order      []oastypes.OrderOption
	inters     []Interceptor
	predicates []predicate.OASTypes
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OASTypesQuery builder.
func (otq *OASTypesQuery) Where(ps ...predicate.OASTypes) *OASTypesQuery {
	otq.predicates = append(otq.predicates, ps...)
	return otq
}

// Limit the number of records to be returned by this query.
func (otq *OASTypesQuery) Limit(limit int) *OASTypesQuery {
	otq.ctx.Limit = &limit
	return otq
}

// Offset to start from.
func (otq *OASTypesQuery) Offset(offset int) *OASTypesQuery {
	otq.ctx.Offset = &offset
	return otq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (otq *OASTypesQuery) Unique(unique bool) *OASTypesQuery {
	otq.ctx.Unique = &unique
	return otq
}

// Order specifies how the records should be ordered.
func (otq *OASTypesQuery) Order(o ...oastypes.OrderOption) *OASTypesQuery {
	otq.order = append(otq.order, o...)
	return otq
}

// First returns the first OASTypes entity from the query.
// Returns a *NotFoundError when no OASTypes was found.
func (otq *OASTypesQuery) First(ctx context.Context) (*OASTypes, error) {
	nodes, err := otq.Limit(1).All(setContextOp(ctx, otq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{oastypes.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (otq *OASTypesQuery) FirstX(ctx context.Context) *OASTypes {
	node, err := otq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OASTypes ID from the query.
// Returns a *NotFoundError when no OASTypes ID was found.
func (otq *OASTypesQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(1).IDs(setContextOp(ctx, otq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{oastypes.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (otq *OASTypesQuery) FirstIDX(ctx context.Context) int {
	id, err := otq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OASTypes entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OASTypes entity is found.
// Returns a *NotFoundError when no OASTypes entities are found.
func (otq *OASTypesQuery) Only(ctx context.Context) (*OASTypes, error) {
	nodes, err := otq.Limit(2).All(setContextOp(ctx, otq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{oastypes.Label}
	default:
		return nil, &NotSingularError{oastypes.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (otq *OASTypesQuery) OnlyX(ctx context.Context) *OASTypes {
	node, err := otq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OASTypes ID in the query.
// Returns a *NotSingularError when more than one OASTypes ID is found.
// Returns a *NotFoundError when no entities are found.
func (otq *OASTypesQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = otq.Limit(2).IDs(setContextOp(ctx, otq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{oastypes.Label}
	default:
		err = &NotSingularError{oastypes.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (otq *OASTypesQuery) OnlyIDX(ctx context.Context) int {
	id, err := otq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OASTypesSlice.
func (otq *OASTypesQuery) All(ctx context.Context) ([]*OASTypes, error) {
	ctx = setContextOp(ctx, otq.ctx, "All")
	if err := otq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OASTypes, *OASTypesQuery]()
	return withInterceptors[[]*OASTypes](ctx, otq, qr, otq.inters)
}

// AllX is like All, but panics if an error occurs.
func (otq *OASTypesQuery) AllX(ctx context.Context) []*OASTypes {
	nodes, err := otq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OASTypes IDs.
func (otq *OASTypesQuery) IDs(ctx context.Context) (ids []int, err error) {
	if otq.ctx.Unique == nil && otq.path != nil {
		otq.Unique(true)
	}
	ctx = setContextOp(ctx, otq.ctx, "IDs")
	if err = otq.Select(oastypes.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (otq *OASTypesQuery) IDsX(ctx context.Context) []int {
	ids, err := otq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (otq *OASTypesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, otq.ctx, "Count")
	if err := otq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, otq, querierCount[*OASTypesQuery](), otq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (otq *OASTypesQuery) CountX(ctx context.Context) int {
	count, err := otq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (otq *OASTypesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, otq.ctx, "Exist")
	switch _, err := otq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("oastypes: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (otq *OASTypesQuery) ExistX(ctx context.Context) bool {
	exist, err := otq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OASTypesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (otq *OASTypesQuery) Clone() *OASTypesQuery {
	if otq == nil {
		return nil
	}
	return &OASTypesQuery{
		config:     otq.config,
		ctx:        otq.ctx.Clone(),
		order:      append([]oastypes.OrderOption{}, otq.order...),
		inters:     append([]Interceptor{}, otq.inters...),
		predicates: append([]predicate.OASTypes{}, otq.predicates...),
		// clone intermediate query.
		sql:  otq.sql.Clone(),
		path: otq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OASTypes.Query().
//		GroupBy(oastypes.FieldInt).
//		Aggregate(oastypes.Count()).
//		Scan(ctx, &v)
func (otq *OASTypesQuery) GroupBy(field string, fields ...string) *OASTypesGroupBy {
	otq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OASTypesGroupBy{build: otq}
	grbuild.flds = &otq.ctx.Fields
	grbuild.label = oastypes.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.OASTypes.Query().
//		Select(oastypes.FieldInt).
//		Scan(ctx, &v)
func (otq *OASTypesQuery) Select(fields ...string) *OASTypesSelect {
	otq.ctx.Fields = append(otq.ctx.Fields, fields...)
	sbuild := &OASTypesSelect{OASTypesQuery: otq}
	sbuild.label = oastypes.Label
	sbuild.flds, sbuild.scan = &otq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OASTypesSelect configured with the given aggregations.
func (otq *OASTypesQuery) Aggregate(fns ...AggregateFunc) *OASTypesSelect {
	return otq.Select().Aggregate(fns...)
}

func (otq *OASTypesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range otq.inters {
		if inter == nil {
			return fmt.Errorf("oastypes: uninitialized interceptor (forgotten import oastypes/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, otq); err != nil {
				return err
			}
		}
	}
	for _, f := range otq.ctx.Fields {
		if !oastypes.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("oastypes: invalid field %q for query", f)}
		}
	}
	if otq.path != nil {
		prev, err := otq.path(ctx)
		if err != nil {
			return err
		}
		otq.sql = prev
	}
	return nil
}

func (otq *OASTypesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OASTypes, error) {
	var (
		nodes = []*OASTypes{}
		_spec = otq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OASTypes).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OASTypes{config: otq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, otq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (otq *OASTypesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := otq.querySpec()
	_spec.Node.Columns = otq.ctx.Fields
	if len(otq.ctx.Fields) > 0 {
		_spec.Unique = otq.ctx.Unique != nil && *otq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, otq.driver, _spec)
}

func (otq *OASTypesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(oastypes.Table, oastypes.Columns, sqlgraph.NewFieldSpec(oastypes.FieldID, field.TypeInt))
	_spec.From = otq.sql
	if unique := otq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if otq.path != nil {
		_spec.Unique = true
	}
	if fields := otq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, oastypes.FieldID)
		for i := range fields {
			if fields[i] != oastypes.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := otq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := otq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := otq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := otq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (otq *OASTypesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(otq.driver.Dialect())
	t1 := builder.Table(oastypes.Table)
	columns := otq.ctx.Fields
	if len(columns) == 0 {
		columns = oastypes.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if otq.sql != nil {
		selector = otq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if otq.ctx.Unique != nil && *otq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range otq.predicates {
		p(selector)
	}
	for _, p := range otq.order {
		p(selector)
	}
	if offset := otq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := otq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OASTypesGroupBy is the group-by builder for OASTypes entities.
type OASTypesGroupBy struct {
	selector
	build *OASTypesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (otgb *OASTypesGroupBy) Aggregate(fns ...AggregateFunc) *OASTypesGroupBy {
	otgb.fns = append(otgb.fns, fns...)
	return otgb
}

// Scan applies the selector query and scans the result into the given value.
func (otgb *OASTypesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, otgb.build.ctx, "GroupBy")
	if err := otgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OASTypesQuery, *OASTypesGroupBy](ctx, otgb.build, otgb, otgb.build.inters, v)
}

func (otgb *OASTypesGroupBy) sqlScan(ctx context.Context, root *OASTypesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(otgb.fns))
	for _, fn := range otgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*otgb.flds)+len(otgb.fns))
		for _, f := range *otgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*otgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := otgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OASTypesSelect is the builder for selecting fields of OASTypes entities.
type OASTypesSelect struct {
	*OASTypesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ots *OASTypesSelect) Aggregate(fns ...AggregateFunc) *OASTypesSelect {
	ots.fns = append(ots.fns, fns...)
	return ots
}

// Scan applies the selector query and scans the result into the given value.
func (ots *OASTypesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ots.ctx, "Select")
	if err := ots.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OASTypesQuery, *OASTypesSelect](ctx, ots.OASTypesQuery, ots, ots.inters, v)
}

func (ots *OASTypesSelect) sqlScan(ctx context.Context, root *OASTypesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ots.fns))
	for _, fn := range ots.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ots.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ots.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
