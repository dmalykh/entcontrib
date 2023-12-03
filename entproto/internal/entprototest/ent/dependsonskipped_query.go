// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/dependsonskipped"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/implicitskippedmessage"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// DependsOnSkippedQuery is the builder for querying DependsOnSkipped entities.
type DependsOnSkippedQuery struct {
	config
	ctx         *QueryContext
	order       []dependsonskipped.OrderOption
	inters      []Interceptor
	predicates  []predicate.DependsOnSkipped
	withSkipped *ImplicitSkippedMessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DependsOnSkippedQuery builder.
func (dosq *DependsOnSkippedQuery) Where(ps ...predicate.DependsOnSkipped) *DependsOnSkippedQuery {
	dosq.predicates = append(dosq.predicates, ps...)
	return dosq
}

// Limit the number of records to be returned by this query.
func (dosq *DependsOnSkippedQuery) Limit(limit int) *DependsOnSkippedQuery {
	dosq.ctx.Limit = &limit
	return dosq
}

// Offset to start from.
func (dosq *DependsOnSkippedQuery) Offset(offset int) *DependsOnSkippedQuery {
	dosq.ctx.Offset = &offset
	return dosq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dosq *DependsOnSkippedQuery) Unique(unique bool) *DependsOnSkippedQuery {
	dosq.ctx.Unique = &unique
	return dosq
}

// Order specifies how the records should be ordered.
func (dosq *DependsOnSkippedQuery) Order(o ...dependsonskipped.OrderOption) *DependsOnSkippedQuery {
	dosq.order = append(dosq.order, o...)
	return dosq
}

// QuerySkipped chains the current query on the "skipped" edge.
func (dosq *DependsOnSkippedQuery) QuerySkipped() *ImplicitSkippedMessageQuery {
	query := (&ImplicitSkippedMessageClient{config: dosq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dosq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dosq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(dependsonskipped.Table, dependsonskipped.FieldID, selector),
			sqlgraph.To(implicitskippedmessage.Table, implicitskippedmessage.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, dependsonskipped.SkippedTable, dependsonskipped.SkippedColumn),
		)
		fromU = sqlgraph.SetNeighbors(dosq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first DependsOnSkipped entity from the query.
// Returns a *NotFoundError when no DependsOnSkipped was found.
func (dosq *DependsOnSkippedQuery) First(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(1).All(setContextOp(ctx, dosq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{dependsonskipped.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DependsOnSkipped ID from the query.
// Returns a *NotFoundError when no DependsOnSkipped ID was found.
func (dosq *DependsOnSkippedQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(1).IDs(setContextOp(ctx, dosq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{dependsonskipped.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) FirstIDX(ctx context.Context) int {
	id, err := dosq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DependsOnSkipped entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DependsOnSkipped entity is found.
// Returns a *NotFoundError when no DependsOnSkipped entities are found.
func (dosq *DependsOnSkippedQuery) Only(ctx context.Context) (*DependsOnSkipped, error) {
	nodes, err := dosq.Limit(2).All(setContextOp(ctx, dosq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{dependsonskipped.Label}
	default:
		return nil, &NotSingularError{dependsonskipped.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyX(ctx context.Context) *DependsOnSkipped {
	node, err := dosq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DependsOnSkipped ID in the query.
// Returns a *NotSingularError when more than one DependsOnSkipped ID is found.
// Returns a *NotFoundError when no entities are found.
func (dosq *DependsOnSkippedQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dosq.Limit(2).IDs(setContextOp(ctx, dosq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{dependsonskipped.Label}
	default:
		err = &NotSingularError{dependsonskipped.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) OnlyIDX(ctx context.Context) int {
	id, err := dosq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DependsOnSkippeds.
func (dosq *DependsOnSkippedQuery) All(ctx context.Context) ([]*DependsOnSkipped, error) {
	ctx = setContextOp(ctx, dosq.ctx, "All")
	if err := dosq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*DependsOnSkipped, *DependsOnSkippedQuery]()
	return withInterceptors[[]*DependsOnSkipped](ctx, dosq, qr, dosq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) AllX(ctx context.Context) []*DependsOnSkipped {
	nodes, err := dosq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DependsOnSkipped IDs.
func (dosq *DependsOnSkippedQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dosq.ctx.Unique == nil && dosq.path != nil {
		dosq.Unique(true)
	}
	ctx = setContextOp(ctx, dosq.ctx, "IDs")
	if err = dosq.Select(dependsonskipped.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) IDsX(ctx context.Context) []int {
	ids, err := dosq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dosq *DependsOnSkippedQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dosq.ctx, "Count")
	if err := dosq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dosq, querierCount[*DependsOnSkippedQuery](), dosq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) CountX(ctx context.Context) int {
	count, err := dosq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dosq *DependsOnSkippedQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dosq.ctx, "Exist")
	switch _, err := dosq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dosq *DependsOnSkippedQuery) ExistX(ctx context.Context) bool {
	exist, err := dosq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DependsOnSkippedQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dosq *DependsOnSkippedQuery) Clone() *DependsOnSkippedQuery {
	if dosq == nil {
		return nil
	}
	return &DependsOnSkippedQuery{
		config:      dosq.config,
		ctx:         dosq.ctx.Clone(),
		order:       append([]dependsonskipped.OrderOption{}, dosq.order...),
		inters:      append([]Interceptor{}, dosq.inters...),
		predicates:  append([]predicate.DependsOnSkipped{}, dosq.predicates...),
		withSkipped: dosq.withSkipped.Clone(),
		// clone intermediate query.
		sql:  dosq.sql.Clone(),
		path: dosq.path,
	}
}

// WithSkipped tells the query-builder to eager-load the nodes that are connected to
// the "skipped" edge. The optional arguments are used to configure the query builder of the edge.
func (dosq *DependsOnSkippedQuery) WithSkipped(opts ...func(*ImplicitSkippedMessageQuery)) *DependsOnSkippedQuery {
	query := (&ImplicitSkippedMessageClient{config: dosq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dosq.withSkipped = query
	return dosq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.DependsOnSkipped.Query().
//		GroupBy(dependsonskipped.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dosq *DependsOnSkippedQuery) GroupBy(field string, fields ...string) *DependsOnSkippedGroupBy {
	dosq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DependsOnSkippedGroupBy{build: dosq}
	grbuild.flds = &dosq.ctx.Fields
	grbuild.label = dependsonskipped.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.DependsOnSkipped.Query().
//		Select(dependsonskipped.FieldName).
//		Scan(ctx, &v)
func (dosq *DependsOnSkippedQuery) Select(fields ...string) *DependsOnSkippedSelect {
	dosq.ctx.Fields = append(dosq.ctx.Fields, fields...)
	sbuild := &DependsOnSkippedSelect{DependsOnSkippedQuery: dosq}
	sbuild.label = dependsonskipped.Label
	sbuild.flds, sbuild.scan = &dosq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DependsOnSkippedSelect configured with the given aggregations.
func (dosq *DependsOnSkippedQuery) Aggregate(fns ...AggregateFunc) *DependsOnSkippedSelect {
	return dosq.Select().Aggregate(fns...)
}

func (dosq *DependsOnSkippedQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dosq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dosq); err != nil {
				return err
			}
		}
	}
	for _, f := range dosq.ctx.Fields {
		if !dependsonskipped.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dosq.path != nil {
		prev, err := dosq.path(ctx)
		if err != nil {
			return err
		}
		dosq.sql = prev
	}
	return nil
}

func (dosq *DependsOnSkippedQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DependsOnSkipped, error) {
	var (
		nodes       = []*DependsOnSkipped{}
		_spec       = dosq.querySpec()
		loadedTypes = [1]bool{
			dosq.withSkipped != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*DependsOnSkipped).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &DependsOnSkipped{config: dosq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dosq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dosq.withSkipped; query != nil {
		if err := dosq.loadSkipped(ctx, query, nodes,
			func(n *DependsOnSkipped) { n.Edges.Skipped = []*ImplicitSkippedMessage{} },
			func(n *DependsOnSkipped, e *ImplicitSkippedMessage) { n.Edges.Skipped = append(n.Edges.Skipped, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dosq *DependsOnSkippedQuery) loadSkipped(ctx context.Context, query *ImplicitSkippedMessageQuery, nodes []*DependsOnSkipped, init func(*DependsOnSkipped), assign func(*DependsOnSkipped, *ImplicitSkippedMessage)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*DependsOnSkipped)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ImplicitSkippedMessage(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(dependsonskipped.SkippedColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.depends_on_skipped_skipped
		if fk == nil {
			return fmt.Errorf(`foreign-key "depends_on_skipped_skipped" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "depends_on_skipped_skipped" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dosq *DependsOnSkippedQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dosq.querySpec()
	_spec.Node.Columns = dosq.ctx.Fields
	if len(dosq.ctx.Fields) > 0 {
		_spec.Unique = dosq.ctx.Unique != nil && *dosq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dosq.driver, _spec)
}

func (dosq *DependsOnSkippedQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(dependsonskipped.Table, dependsonskipped.Columns, sqlgraph.NewFieldSpec(dependsonskipped.FieldID, field.TypeInt))
	_spec.From = dosq.sql
	if unique := dosq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dosq.path != nil {
		_spec.Unique = true
	}
	if fields := dosq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dependsonskipped.FieldID)
		for i := range fields {
			if fields[i] != dependsonskipped.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dosq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dosq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dosq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dosq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dosq *DependsOnSkippedQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dosq.driver.Dialect())
	t1 := builder.Table(dependsonskipped.Table)
	columns := dosq.ctx.Fields
	if len(columns) == 0 {
		columns = dependsonskipped.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dosq.sql != nil {
		selector = dosq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dosq.ctx.Unique != nil && *dosq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dosq.predicates {
		p(selector)
	}
	for _, p := range dosq.order {
		p(selector)
	}
	if offset := dosq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dosq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DependsOnSkippedGroupBy is the group-by builder for DependsOnSkipped entities.
type DependsOnSkippedGroupBy struct {
	selector
	build *DependsOnSkippedQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dosgb *DependsOnSkippedGroupBy) Aggregate(fns ...AggregateFunc) *DependsOnSkippedGroupBy {
	dosgb.fns = append(dosgb.fns, fns...)
	return dosgb
}

// Scan applies the selector query and scans the result into the given value.
func (dosgb *DependsOnSkippedGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dosgb.build.ctx, "GroupBy")
	if err := dosgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DependsOnSkippedQuery, *DependsOnSkippedGroupBy](ctx, dosgb.build, dosgb, dosgb.build.inters, v)
}

func (dosgb *DependsOnSkippedGroupBy) sqlScan(ctx context.Context, root *DependsOnSkippedQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dosgb.fns))
	for _, fn := range dosgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dosgb.flds)+len(dosgb.fns))
		for _, f := range *dosgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dosgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dosgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DependsOnSkippedSelect is the builder for selecting fields of DependsOnSkipped entities.
type DependsOnSkippedSelect struct {
	*DependsOnSkippedQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (doss *DependsOnSkippedSelect) Aggregate(fns ...AggregateFunc) *DependsOnSkippedSelect {
	doss.fns = append(doss.fns, fns...)
	return doss
}

// Scan applies the selector query and scans the result into the given value.
func (doss *DependsOnSkippedSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, doss.ctx, "Select")
	if err := doss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DependsOnSkippedQuery, *DependsOnSkippedSelect](ctx, doss.DependsOnSkippedQuery, doss, doss.inters, v)
}

func (doss *DependsOnSkippedSelect) sqlScan(ctx context.Context, root *DependsOnSkippedQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(doss.fns))
	for _, fn := range doss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*doss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := doss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
