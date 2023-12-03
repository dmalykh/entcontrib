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
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/image"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/nobackref"
	"github.com/dmalykh/entcontrib/entproto/internal/entprototest/ent/predicate"
)

// NoBackrefQuery is the builder for querying NoBackref entities.
type NoBackrefQuery struct {
	config
	ctx        *QueryContext
	order      []nobackref.OrderOption
	inters     []Interceptor
	predicates []predicate.NoBackref
	withImages *ImageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NoBackrefQuery builder.
func (nbq *NoBackrefQuery) Where(ps ...predicate.NoBackref) *NoBackrefQuery {
	nbq.predicates = append(nbq.predicates, ps...)
	return nbq
}

// Limit the number of records to be returned by this query.
func (nbq *NoBackrefQuery) Limit(limit int) *NoBackrefQuery {
	nbq.ctx.Limit = &limit
	return nbq
}

// Offset to start from.
func (nbq *NoBackrefQuery) Offset(offset int) *NoBackrefQuery {
	nbq.ctx.Offset = &offset
	return nbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nbq *NoBackrefQuery) Unique(unique bool) *NoBackrefQuery {
	nbq.ctx.Unique = &unique
	return nbq
}

// Order specifies how the records should be ordered.
func (nbq *NoBackrefQuery) Order(o ...nobackref.OrderOption) *NoBackrefQuery {
	nbq.order = append(nbq.order, o...)
	return nbq
}

// QueryImages chains the current query on the "images" edge.
func (nbq *NoBackrefQuery) QueryImages() *ImageQuery {
	query := (&ImageClient{config: nbq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nbq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(nobackref.Table, nobackref.FieldID, selector),
			sqlgraph.To(image.Table, image.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, nobackref.ImagesTable, nobackref.ImagesColumn),
		)
		fromU = sqlgraph.SetNeighbors(nbq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NoBackref entity from the query.
// Returns a *NotFoundError when no NoBackref was found.
func (nbq *NoBackrefQuery) First(ctx context.Context) (*NoBackref, error) {
	nodes, err := nbq.Limit(1).All(setContextOp(ctx, nbq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{nobackref.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nbq *NoBackrefQuery) FirstX(ctx context.Context) *NoBackref {
	node, err := nbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NoBackref ID from the query.
// Returns a *NotFoundError when no NoBackref ID was found.
func (nbq *NoBackrefQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nbq.Limit(1).IDs(setContextOp(ctx, nbq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{nobackref.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nbq *NoBackrefQuery) FirstIDX(ctx context.Context) int {
	id, err := nbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NoBackref entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NoBackref entity is found.
// Returns a *NotFoundError when no NoBackref entities are found.
func (nbq *NoBackrefQuery) Only(ctx context.Context) (*NoBackref, error) {
	nodes, err := nbq.Limit(2).All(setContextOp(ctx, nbq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{nobackref.Label}
	default:
		return nil, &NotSingularError{nobackref.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nbq *NoBackrefQuery) OnlyX(ctx context.Context) *NoBackref {
	node, err := nbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NoBackref ID in the query.
// Returns a *NotSingularError when more than one NoBackref ID is found.
// Returns a *NotFoundError when no entities are found.
func (nbq *NoBackrefQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = nbq.Limit(2).IDs(setContextOp(ctx, nbq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{nobackref.Label}
	default:
		err = &NotSingularError{nobackref.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nbq *NoBackrefQuery) OnlyIDX(ctx context.Context) int {
	id, err := nbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NoBackrefs.
func (nbq *NoBackrefQuery) All(ctx context.Context) ([]*NoBackref, error) {
	ctx = setContextOp(ctx, nbq.ctx, "All")
	if err := nbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NoBackref, *NoBackrefQuery]()
	return withInterceptors[[]*NoBackref](ctx, nbq, qr, nbq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nbq *NoBackrefQuery) AllX(ctx context.Context) []*NoBackref {
	nodes, err := nbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NoBackref IDs.
func (nbq *NoBackrefQuery) IDs(ctx context.Context) (ids []int, err error) {
	if nbq.ctx.Unique == nil && nbq.path != nil {
		nbq.Unique(true)
	}
	ctx = setContextOp(ctx, nbq.ctx, "IDs")
	if err = nbq.Select(nobackref.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nbq *NoBackrefQuery) IDsX(ctx context.Context) []int {
	ids, err := nbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nbq *NoBackrefQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nbq.ctx, "Count")
	if err := nbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nbq, querierCount[*NoBackrefQuery](), nbq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nbq *NoBackrefQuery) CountX(ctx context.Context) int {
	count, err := nbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nbq *NoBackrefQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nbq.ctx, "Exist")
	switch _, err := nbq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nbq *NoBackrefQuery) ExistX(ctx context.Context) bool {
	exist, err := nbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NoBackrefQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nbq *NoBackrefQuery) Clone() *NoBackrefQuery {
	if nbq == nil {
		return nil
	}
	return &NoBackrefQuery{
		config:     nbq.config,
		ctx:        nbq.ctx.Clone(),
		order:      append([]nobackref.OrderOption{}, nbq.order...),
		inters:     append([]Interceptor{}, nbq.inters...),
		predicates: append([]predicate.NoBackref{}, nbq.predicates...),
		withImages: nbq.withImages.Clone(),
		// clone intermediate query.
		sql:  nbq.sql.Clone(),
		path: nbq.path,
	}
}

// WithImages tells the query-builder to eager-load the nodes that are connected to
// the "images" edge. The optional arguments are used to configure the query builder of the edge.
func (nbq *NoBackrefQuery) WithImages(opts ...func(*ImageQuery)) *NoBackrefQuery {
	query := (&ImageClient{config: nbq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nbq.withImages = query
	return nbq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (nbq *NoBackrefQuery) GroupBy(field string, fields ...string) *NoBackrefGroupBy {
	nbq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NoBackrefGroupBy{build: nbq}
	grbuild.flds = &nbq.ctx.Fields
	grbuild.label = nobackref.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (nbq *NoBackrefQuery) Select(fields ...string) *NoBackrefSelect {
	nbq.ctx.Fields = append(nbq.ctx.Fields, fields...)
	sbuild := &NoBackrefSelect{NoBackrefQuery: nbq}
	sbuild.label = nobackref.Label
	sbuild.flds, sbuild.scan = &nbq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NoBackrefSelect configured with the given aggregations.
func (nbq *NoBackrefQuery) Aggregate(fns ...AggregateFunc) *NoBackrefSelect {
	return nbq.Select().Aggregate(fns...)
}

func (nbq *NoBackrefQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nbq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nbq); err != nil {
				return err
			}
		}
	}
	for _, f := range nbq.ctx.Fields {
		if !nobackref.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nbq.path != nil {
		prev, err := nbq.path(ctx)
		if err != nil {
			return err
		}
		nbq.sql = prev
	}
	return nil
}

func (nbq *NoBackrefQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NoBackref, error) {
	var (
		nodes       = []*NoBackref{}
		_spec       = nbq.querySpec()
		loadedTypes = [1]bool{
			nbq.withImages != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NoBackref).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NoBackref{config: nbq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nbq.withImages; query != nil {
		if err := nbq.loadImages(ctx, query, nodes,
			func(n *NoBackref) { n.Edges.Images = []*Image{} },
			func(n *NoBackref, e *Image) { n.Edges.Images = append(n.Edges.Images, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nbq *NoBackrefQuery) loadImages(ctx context.Context, query *ImageQuery, nodes []*NoBackref, init func(*NoBackref), assign func(*NoBackref, *Image)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*NoBackref)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Image(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(nobackref.ImagesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.no_backref_images
		if fk == nil {
			return fmt.Errorf(`foreign-key "no_backref_images" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "no_backref_images" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (nbq *NoBackrefQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nbq.querySpec()
	_spec.Node.Columns = nbq.ctx.Fields
	if len(nbq.ctx.Fields) > 0 {
		_spec.Unique = nbq.ctx.Unique != nil && *nbq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nbq.driver, _spec)
}

func (nbq *NoBackrefQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(nobackref.Table, nobackref.Columns, sqlgraph.NewFieldSpec(nobackref.FieldID, field.TypeInt))
	_spec.From = nbq.sql
	if unique := nbq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nbq.path != nil {
		_spec.Unique = true
	}
	if fields := nbq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nobackref.FieldID)
		for i := range fields {
			if fields[i] != nobackref.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nbq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nbq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nbq *NoBackrefQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nbq.driver.Dialect())
	t1 := builder.Table(nobackref.Table)
	columns := nbq.ctx.Fields
	if len(columns) == 0 {
		columns = nobackref.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nbq.sql != nil {
		selector = nbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nbq.ctx.Unique != nil && *nbq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nbq.predicates {
		p(selector)
	}
	for _, p := range nbq.order {
		p(selector)
	}
	if offset := nbq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nbq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NoBackrefGroupBy is the group-by builder for NoBackref entities.
type NoBackrefGroupBy struct {
	selector
	build *NoBackrefQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (nbgb *NoBackrefGroupBy) Aggregate(fns ...AggregateFunc) *NoBackrefGroupBy {
	nbgb.fns = append(nbgb.fns, fns...)
	return nbgb
}

// Scan applies the selector query and scans the result into the given value.
func (nbgb *NoBackrefGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nbgb.build.ctx, "GroupBy")
	if err := nbgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NoBackrefQuery, *NoBackrefGroupBy](ctx, nbgb.build, nbgb, nbgb.build.inters, v)
}

func (nbgb *NoBackrefGroupBy) sqlScan(ctx context.Context, root *NoBackrefQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(nbgb.fns))
	for _, fn := range nbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*nbgb.flds)+len(nbgb.fns))
		for _, f := range *nbgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*nbgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nbgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NoBackrefSelect is the builder for selecting fields of NoBackref entities.
type NoBackrefSelect struct {
	*NoBackrefQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nbs *NoBackrefSelect) Aggregate(fns ...AggregateFunc) *NoBackrefSelect {
	nbs.fns = append(nbs.fns, fns...)
	return nbs
}

// Scan applies the selector query and scans the result into the given value.
func (nbs *NoBackrefSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nbs.ctx, "Select")
	if err := nbs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NoBackrefQuery, *NoBackrefSelect](ctx, nbs.NoBackrefQuery, nbs, nbs.inters, v)
}

func (nbs *NoBackrefSelect) sqlScan(ctx context.Context, root *NoBackrefQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nbs.fns))
	for _, fn := range nbs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nbs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
