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
	"github.com/google/uuid"
	"github.com/scorify/scorify/pkg/ent/inject"
	"github.com/scorify/scorify/pkg/ent/injectsubmission"
	"github.com/scorify/scorify/pkg/ent/predicate"
)

// InjectQuery is the builder for querying Inject entities.
type InjectQuery struct {
	config
	ctx             *QueryContext
	order           []inject.OrderOption
	inters          []Interceptor
	predicates      []predicate.Inject
	withSubmissions *InjectSubmissionQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InjectQuery builder.
func (iq *InjectQuery) Where(ps ...predicate.Inject) *InjectQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InjectQuery) Limit(limit int) *InjectQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InjectQuery) Offset(offset int) *InjectQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InjectQuery) Unique(unique bool) *InjectQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InjectQuery) Order(o ...inject.OrderOption) *InjectQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// QuerySubmissions chains the current query on the "submissions" edge.
func (iq *InjectQuery) QuerySubmissions() *InjectSubmissionQuery {
	query := (&InjectSubmissionClient{config: iq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := iq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := iq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(inject.Table, inject.FieldID, selector),
			sqlgraph.To(injectsubmission.Table, injectsubmission.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, inject.SubmissionsTable, inject.SubmissionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(iq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Inject entity from the query.
// Returns a *NotFoundError when no Inject was found.
func (iq *InjectQuery) First(ctx context.Context) (*Inject, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InjectQuery) FirstX(ctx context.Context) *Inject {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Inject ID from the query.
// Returns a *NotFoundError when no Inject ID was found.
func (iq *InjectQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InjectQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Inject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Inject entity is found.
// Returns a *NotFoundError when no Inject entities are found.
func (iq *InjectQuery) Only(ctx context.Context) (*Inject, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inject.Label}
	default:
		return nil, &NotSingularError{inject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InjectQuery) OnlyX(ctx context.Context) *Inject {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Inject ID in the query.
// Returns a *NotSingularError when more than one Inject ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InjectQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inject.Label}
	default:
		err = &NotSingularError{inject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InjectQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Injects.
func (iq *InjectQuery) All(ctx context.Context) ([]*Inject, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Inject, *InjectQuery]()
	return withInterceptors[[]*Inject](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InjectQuery) AllX(ctx context.Context) []*Inject {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Inject IDs.
func (iq *InjectQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(inject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InjectQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InjectQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InjectQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InjectQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InjectQuery) Clone() *InjectQuery {
	if iq == nil {
		return nil
	}
	return &InjectQuery{
		config:          iq.config,
		ctx:             iq.ctx.Clone(),
		order:           append([]inject.OrderOption{}, iq.order...),
		inters:          append([]Interceptor{}, iq.inters...),
		predicates:      append([]predicate.Inject{}, iq.predicates...),
		withSubmissions: iq.withSubmissions.Clone(),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
	}
}

// WithSubmissions tells the query-builder to eager-load the nodes that are connected to
// the "submissions" edge. The optional arguments are used to configure the query builder of the edge.
func (iq *InjectQuery) WithSubmissions(opts ...func(*InjectSubmissionQuery)) *InjectQuery {
	query := (&InjectSubmissionClient{config: iq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	iq.withSubmissions = query
	return iq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Inject.Query().
//		GroupBy(inject.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InjectQuery) GroupBy(field string, fields ...string) *InjectGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InjectGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = inject.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Inject.Query().
//		Select(inject.FieldCreateTime).
//		Scan(ctx, &v)
func (iq *InjectQuery) Select(fields ...string) *InjectSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InjectSelect{InjectQuery: iq}
	sbuild.label = inject.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InjectSelect configured with the given aggregations.
func (iq *InjectQuery) Aggregate(fns ...AggregateFunc) *InjectSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !inject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Inject, error) {
	var (
		nodes       = []*Inject{}
		_spec       = iq.querySpec()
		loadedTypes = [1]bool{
			iq.withSubmissions != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Inject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Inject{config: iq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := iq.withSubmissions; query != nil {
		if err := iq.loadSubmissions(ctx, query, nodes,
			func(n *Inject) { n.Edges.Submissions = []*InjectSubmission{} },
			func(n *Inject, e *InjectSubmission) { n.Edges.Submissions = append(n.Edges.Submissions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (iq *InjectQuery) loadSubmissions(ctx context.Context, query *InjectSubmissionQuery, nodes []*Inject, init func(*Inject), assign func(*Inject, *InjectSubmission)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Inject)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(injectsubmission.FieldInjectID)
	}
	query.Where(predicate.InjectSubmission(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(inject.SubmissionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.InjectID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "inject_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (iq *InjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(inject.Table, inject.Columns, sqlgraph.NewFieldSpec(inject.FieldID, field.TypeUUID))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inject.FieldID)
		for i := range fields {
			if fields[i] != inject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(inject.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = inject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InjectGroupBy is the group-by builder for Inject entities.
type InjectGroupBy struct {
	selector
	build *InjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InjectGroupBy) Aggregate(fns ...AggregateFunc) *InjectGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InjectQuery, *InjectGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InjectGroupBy) sqlScan(ctx context.Context, root *InjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InjectSelect is the builder for selecting fields of Inject entities.
type InjectSelect struct {
	*InjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InjectSelect) Aggregate(fns ...AggregateFunc) *InjectSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InjectQuery, *InjectSelect](ctx, is.InjectQuery, is, is.inters, v)
}

func (is *InjectSelect) sqlScan(ctx context.Context, root *InjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
