// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/twomethodservice"
)

// TwoMethodServiceQuery is the builder for querying TwoMethodService entities.
type TwoMethodServiceQuery struct {
	config
	ctx        *QueryContext
	order      []twomethodservice.OrderOption
	inters     []Interceptor
	predicates []predicate.TwoMethodService
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwoMethodServiceQuery builder.
func (tmsq *TwoMethodServiceQuery) Where(ps ...predicate.TwoMethodService) *TwoMethodServiceQuery {
	tmsq.predicates = append(tmsq.predicates, ps...)
	return tmsq
}

// Limit the number of records to be returned by this query.
func (tmsq *TwoMethodServiceQuery) Limit(limit int) *TwoMethodServiceQuery {
	tmsq.ctx.Limit = &limit
	return tmsq
}

// Offset to start from.
func (tmsq *TwoMethodServiceQuery) Offset(offset int) *TwoMethodServiceQuery {
	tmsq.ctx.Offset = &offset
	return tmsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmsq *TwoMethodServiceQuery) Unique(unique bool) *TwoMethodServiceQuery {
	tmsq.ctx.Unique = &unique
	return tmsq
}

// Order specifies how the records should be ordered.
func (tmsq *TwoMethodServiceQuery) Order(o ...twomethodservice.OrderOption) *TwoMethodServiceQuery {
	tmsq.order = append(tmsq.order, o...)
	return tmsq
}

// First returns the first TwoMethodService entity from the query.
// Returns a *NotFoundError when no TwoMethodService was found.
func (tmsq *TwoMethodServiceQuery) First(ctx context.Context) (*TwoMethodService, error) {
	nodes, err := tmsq.Limit(1).All(setContextOp(ctx, tmsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twomethodservice.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) FirstX(ctx context.Context) *TwoMethodService {
	node, err := tmsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwoMethodService ID from the query.
// Returns a *NotFoundError when no TwoMethodService ID was found.
func (tmsq *TwoMethodServiceQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tmsq.Limit(1).IDs(setContextOp(ctx, tmsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twomethodservice.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) FirstIDX(ctx context.Context) int {
	id, err := tmsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwoMethodService entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwoMethodService entity is found.
// Returns a *NotFoundError when no TwoMethodService entities are found.
func (tmsq *TwoMethodServiceQuery) Only(ctx context.Context) (*TwoMethodService, error) {
	nodes, err := tmsq.Limit(2).All(setContextOp(ctx, tmsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twomethodservice.Label}
	default:
		return nil, &NotSingularError{twomethodservice.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) OnlyX(ctx context.Context) *TwoMethodService {
	node, err := tmsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwoMethodService ID in the query.
// Returns a *NotSingularError when more than one TwoMethodService ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmsq *TwoMethodServiceQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tmsq.Limit(2).IDs(setContextOp(ctx, tmsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twomethodservice.Label}
	default:
		err = &NotSingularError{twomethodservice.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) OnlyIDX(ctx context.Context) int {
	id, err := tmsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwoMethodServices.
func (tmsq *TwoMethodServiceQuery) All(ctx context.Context) ([]*TwoMethodService, error) {
	ctx = setContextOp(ctx, tmsq.ctx, "All")
	if err := tmsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwoMethodService, *TwoMethodServiceQuery]()
	return withInterceptors[[]*TwoMethodService](ctx, tmsq, qr, tmsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) AllX(ctx context.Context) []*TwoMethodService {
	nodes, err := tmsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwoMethodService IDs.
func (tmsq *TwoMethodServiceQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tmsq.ctx.Unique == nil && tmsq.path != nil {
		tmsq.Unique(true)
	}
	ctx = setContextOp(ctx, tmsq.ctx, "IDs")
	if err = tmsq.Select(twomethodservice.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) IDsX(ctx context.Context) []int {
	ids, err := tmsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmsq *TwoMethodServiceQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tmsq.ctx, "Count")
	if err := tmsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tmsq, querierCount[*TwoMethodServiceQuery](), tmsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) CountX(ctx context.Context) int {
	count, err := tmsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmsq *TwoMethodServiceQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tmsq.ctx, "Exist")
	switch _, err := tmsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tmsq *TwoMethodServiceQuery) ExistX(ctx context.Context) bool {
	exist, err := tmsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwoMethodServiceQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmsq *TwoMethodServiceQuery) Clone() *TwoMethodServiceQuery {
	if tmsq == nil {
		return nil
	}
	return &TwoMethodServiceQuery{
		config:     tmsq.config,
		ctx:        tmsq.ctx.Clone(),
		order:      append([]twomethodservice.OrderOption{}, tmsq.order...),
		inters:     append([]Interceptor{}, tmsq.inters...),
		predicates: append([]predicate.TwoMethodService{}, tmsq.predicates...),
		// clone intermediate query.
		sql:  tmsq.sql.Clone(),
		path: tmsq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (tmsq *TwoMethodServiceQuery) GroupBy(field string, fields ...string) *TwoMethodServiceGroupBy {
	tmsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwoMethodServiceGroupBy{build: tmsq}
	grbuild.flds = &tmsq.ctx.Fields
	grbuild.label = twomethodservice.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (tmsq *TwoMethodServiceQuery) Select(fields ...string) *TwoMethodServiceSelect {
	tmsq.ctx.Fields = append(tmsq.ctx.Fields, fields...)
	sbuild := &TwoMethodServiceSelect{TwoMethodServiceQuery: tmsq}
	sbuild.label = twomethodservice.Label
	sbuild.flds, sbuild.scan = &tmsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwoMethodServiceSelect configured with the given aggregations.
func (tmsq *TwoMethodServiceQuery) Aggregate(fns ...AggregateFunc) *TwoMethodServiceSelect {
	return tmsq.Select().Aggregate(fns...)
}

func (tmsq *TwoMethodServiceQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tmsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tmsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tmsq.ctx.Fields {
		if !twomethodservice.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tmsq.path != nil {
		prev, err := tmsq.path(ctx)
		if err != nil {
			return err
		}
		tmsq.sql = prev
	}
	return nil
}

func (tmsq *TwoMethodServiceQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwoMethodService, error) {
	var (
		nodes = []*TwoMethodService{}
		_spec = tmsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwoMethodService).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwoMethodService{config: tmsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmsq *TwoMethodServiceQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmsq.querySpec()
	_spec.Node.Columns = tmsq.ctx.Fields
	if len(tmsq.ctx.Fields) > 0 {
		_spec.Unique = tmsq.ctx.Unique != nil && *tmsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tmsq.driver, _spec)
}

func (tmsq *TwoMethodServiceQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twomethodservice.Table, twomethodservice.Columns, sqlgraph.NewFieldSpec(twomethodservice.FieldID, field.TypeInt))
	_spec.From = tmsq.sql
	if unique := tmsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tmsq.path != nil {
		_spec.Unique = true
	}
	if fields := tmsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twomethodservice.FieldID)
		for i := range fields {
			if fields[i] != twomethodservice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmsq *TwoMethodServiceQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmsq.driver.Dialect())
	t1 := builder.Table(twomethodservice.Table)
	columns := tmsq.ctx.Fields
	if len(columns) == 0 {
		columns = twomethodservice.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmsq.sql != nil {
		selector = tmsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmsq.ctx.Unique != nil && *tmsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tmsq.predicates {
		p(selector)
	}
	for _, p := range tmsq.order {
		p(selector)
	}
	if offset := tmsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TwoMethodServiceGroupBy is the group-by builder for TwoMethodService entities.
type TwoMethodServiceGroupBy struct {
	selector
	build *TwoMethodServiceQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmsgb *TwoMethodServiceGroupBy) Aggregate(fns ...AggregateFunc) *TwoMethodServiceGroupBy {
	tmsgb.fns = append(tmsgb.fns, fns...)
	return tmsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tmsgb *TwoMethodServiceGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmsgb.build.ctx, "GroupBy")
	if err := tmsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwoMethodServiceQuery, *TwoMethodServiceGroupBy](ctx, tmsgb.build, tmsgb, tmsgb.build.inters, v)
}

func (tmsgb *TwoMethodServiceGroupBy) sqlScan(ctx context.Context, root *TwoMethodServiceQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tmsgb.fns))
	for _, fn := range tmsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tmsgb.flds)+len(tmsgb.fns))
		for _, f := range *tmsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tmsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwoMethodServiceSelect is the builder for selecting fields of TwoMethodService entities.
type TwoMethodServiceSelect struct {
	*TwoMethodServiceQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tmss *TwoMethodServiceSelect) Aggregate(fns ...AggregateFunc) *TwoMethodServiceSelect {
	tmss.fns = append(tmss.fns, fns...)
	return tmss
}

// Scan applies the selector query and scans the result into the given value.
func (tmss *TwoMethodServiceSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tmss.ctx, "Select")
	if err := tmss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwoMethodServiceQuery, *TwoMethodServiceSelect](ctx, tmss.TwoMethodServiceQuery, tmss, tmss.inters, v)
}

func (tmss *TwoMethodServiceSelect) sqlScan(ctx context.Context, root *TwoMethodServiceQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tmss.fns))
	for _, fn := range tmss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tmss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
