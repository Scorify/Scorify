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
	"github.com/scorify/scorify/pkg/ent/kothstatus"
	"github.com/scorify/scorify/pkg/ent/minion"
	"github.com/scorify/scorify/pkg/ent/predicate"
	"github.com/scorify/scorify/pkg/ent/status"
)

// MinionQuery is the builder for querying Minion entities.
type MinionQuery struct {
	config
	ctx              *QueryContext
	order            []minion.OrderOption
	inters           []Interceptor
	predicates       []predicate.Minion
	withStatuses     *StatusQuery
	withKothStatuses *KothStatusQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MinionQuery builder.
func (mq *MinionQuery) Where(ps ...predicate.Minion) *MinionQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MinionQuery) Limit(limit int) *MinionQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MinionQuery) Offset(offset int) *MinionQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MinionQuery) Unique(unique bool) *MinionQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MinionQuery) Order(o ...minion.OrderOption) *MinionQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryStatuses chains the current query on the "statuses" edge.
func (mq *MinionQuery) QueryStatuses() *StatusQuery {
	query := (&StatusClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(minion.Table, minion.FieldID, selector),
			sqlgraph.To(status.Table, status.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, minion.StatusesTable, minion.StatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryKothStatuses chains the current query on the "kothStatuses" edge.
func (mq *MinionQuery) QueryKothStatuses() *KothStatusQuery {
	query := (&KothStatusClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(minion.Table, minion.FieldID, selector),
			sqlgraph.To(kothstatus.Table, kothstatus.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, minion.KothStatusesTable, minion.KothStatusesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Minion entity from the query.
// Returns a *NotFoundError when no Minion was found.
func (mq *MinionQuery) First(ctx context.Context) (*Minion, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{minion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MinionQuery) FirstX(ctx context.Context) *Minion {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Minion ID from the query.
// Returns a *NotFoundError when no Minion ID was found.
func (mq *MinionQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{minion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MinionQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Minion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Minion entity is found.
// Returns a *NotFoundError when no Minion entities are found.
func (mq *MinionQuery) Only(ctx context.Context) (*Minion, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{minion.Label}
	default:
		return nil, &NotSingularError{minion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MinionQuery) OnlyX(ctx context.Context) *Minion {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Minion ID in the query.
// Returns a *NotSingularError when more than one Minion ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MinionQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{minion.Label}
	default:
		err = &NotSingularError{minion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MinionQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Minions.
func (mq *MinionQuery) All(ctx context.Context) ([]*Minion, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Minion, *MinionQuery]()
	return withInterceptors[[]*Minion](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MinionQuery) AllX(ctx context.Context) []*Minion {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Minion IDs.
func (mq *MinionQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(minion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MinionQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MinionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MinionQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MinionQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MinionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MinionQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MinionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MinionQuery) Clone() *MinionQuery {
	if mq == nil {
		return nil
	}
	return &MinionQuery{
		config:           mq.config,
		ctx:              mq.ctx.Clone(),
		order:            append([]minion.OrderOption{}, mq.order...),
		inters:           append([]Interceptor{}, mq.inters...),
		predicates:       append([]predicate.Minion{}, mq.predicates...),
		withStatuses:     mq.withStatuses.Clone(),
		withKothStatuses: mq.withKothStatuses.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithStatuses tells the query-builder to eager-load the nodes that are connected to
// the "statuses" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MinionQuery) WithStatuses(opts ...func(*StatusQuery)) *MinionQuery {
	query := (&StatusClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withStatuses = query
	return mq
}

// WithKothStatuses tells the query-builder to eager-load the nodes that are connected to
// the "kothStatuses" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MinionQuery) WithKothStatuses(opts ...func(*KothStatusQuery)) *MinionQuery {
	query := (&KothStatusClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withKothStatuses = query
	return mq
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
//	client.Minion.Query().
//		GroupBy(minion.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MinionQuery) GroupBy(field string, fields ...string) *MinionGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MinionGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = minion.Label
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
//	client.Minion.Query().
//		Select(minion.FieldCreateTime).
//		Scan(ctx, &v)
func (mq *MinionQuery) Select(fields ...string) *MinionSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MinionSelect{MinionQuery: mq}
	sbuild.label = minion.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MinionSelect configured with the given aggregations.
func (mq *MinionQuery) Aggregate(fns ...AggregateFunc) *MinionSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MinionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !minion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MinionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Minion, error) {
	var (
		nodes       = []*Minion{}
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withStatuses != nil,
			mq.withKothStatuses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Minion).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Minion{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withStatuses; query != nil {
		if err := mq.loadStatuses(ctx, query, nodes,
			func(n *Minion) { n.Edges.Statuses = []*Status{} },
			func(n *Minion, e *Status) { n.Edges.Statuses = append(n.Edges.Statuses, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withKothStatuses; query != nil {
		if err := mq.loadKothStatuses(ctx, query, nodes,
			func(n *Minion) { n.Edges.KothStatuses = []*KothStatus{} },
			func(n *Minion, e *KothStatus) { n.Edges.KothStatuses = append(n.Edges.KothStatuses, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MinionQuery) loadStatuses(ctx context.Context, query *StatusQuery, nodes []*Minion, init func(*Minion), assign func(*Minion, *Status)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Minion)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(status.FieldMinionID)
	}
	query.Where(predicate.Status(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(minion.StatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MinionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "minion_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mq *MinionQuery) loadKothStatuses(ctx context.Context, query *KothStatusQuery, nodes []*Minion, init func(*Minion), assign func(*Minion, *KothStatus)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*Minion)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(kothstatus.FieldMinionID)
	}
	query.Where(predicate.KothStatus(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(minion.KothStatusesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MinionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "minion_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MinionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MinionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(minion.Table, minion.Columns, sqlgraph.NewFieldSpec(minion.FieldID, field.TypeUUID))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, minion.FieldID)
		for i := range fields {
			if fields[i] != minion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MinionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(minion.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = minion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MinionGroupBy is the group-by builder for Minion entities.
type MinionGroupBy struct {
	selector
	build *MinionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MinionGroupBy) Aggregate(fns ...AggregateFunc) *MinionGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MinionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MinionQuery, *MinionGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MinionGroupBy) sqlScan(ctx context.Context, root *MinionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MinionSelect is the builder for selecting fields of Minion entities.
type MinionSelect struct {
	*MinionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MinionSelect) Aggregate(fns ...AggregateFunc) *MinionSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MinionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MinionQuery, *MinionSelect](ctx, ms.MinionQuery, ms, ms.inters, v)
}

func (ms *MinionSelect) sqlScan(ctx context.Context, root *MinionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
